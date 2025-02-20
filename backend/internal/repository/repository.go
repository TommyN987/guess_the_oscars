package repository

import (
	"context"
	"errors"

	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CheckHealth(ctx context.Context) (string, error)
	CreateUser(ctx context.Context, user domain.User, userValidation domain.UserValidation) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByToken(ctx context.Context, token string) (domain.User, error)
	UpdateUserConfirmation(ctx context.Context, userID int) error
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetNominationsByCategoryID(ctx context.Context, userID, categoryID int) ([]domain.Nomination, *domain.Guess, error)
	UpsertGuess(ctx context.Context, guess domain.Guess) error
}

type PgxRepository struct {
	db *pgxpool.Pool
}

func NewPgxRepository(db *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{db: db}
}

func (r *PgxRepository) CheckHealth(ctx context.Context) (string, error) {
	err := r.db.Ping(ctx)
	if err != nil {
		return "", err
	}

	return "Database connection healthy", nil

}

func (r *PgxRepository) CreateUser(ctx context.Context, user domain.User, userValidation domain.UserValidation) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO users (name, email, password_hash, confirmation_token, confirmation_expires)
        VALUES ($1, $2, $3, $4, $5)`,
		user.Name, user.Email, user.Password, userValidation.ConfirmationToken, userValidation.ConfirmationExpires)
	return err
}

func (r *PgxRepository) UpdateUserConfirmation(ctx context.Context, userID int) error {
	_, err := r.db.Exec(ctx, `
        UPDATE users
        SET email_confirmed = TRUE, confirmation_token = NULL, confirmation_expires = NULL
        WHERE id = $1
        `, userID)
	return err
}

func (r *PgxRepository) GetUserByToken(ctx context.Context, token string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, `
        SELECT id, name, email, password_hash, email_confirmed
        FROM users
        WHERE confirmation_token = $1 AND confirmation_expires > NOW()
        `, token).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.EmailConfirmed,
	)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *PgxRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, `
        SELECT id, name, email, password_hash, email_confirmed
        FROM users
        WHERE email = $1
        `, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.EmailConfirmed)

	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *PgxRepository) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var c CategoryDB
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, toDomainCategory(c))
	}

	return categories, nil
}

func (r *PgxRepository) GetNominationsByCategoryID(ctx context.Context, userID, categoryID int) ([]domain.Nomination, *domain.Guess, error) {
	category, err := r.fetchCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, nil, err
	}

	nominationsDB, err := r.fetchNominationsByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, nil, err
	}

	nominations := make([]domain.Nomination, 0, len(nominationsDB))
	var guess *domain.Guess = nil

	for _, n := range nominationsDB {
		nomination, err := r.buildNomination(ctx, n, category)
		if err != nil {
			return nil, nil, err
		}
		nominations = append(nominations, nomination)

		// Only set the guess if it hasn't been set already
		if guess == nil {
			g, err := r.getGuessByNominationAndUserID(ctx, userID, nomination.ID)
			if err == nil {
				guess = &g
			} else if err != pgx.ErrNoRows {
				return nil, nil, err
			}
		}
	}

	return nominations, guess, nil
}

func (r *PgxRepository) UpsertGuess(ctx context.Context, g domain.Guess) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO guesses (user_id, nomination_id, category_id)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id, category_id) 
        DO UPDATE SET nomination_id = EXCLUDED.nomination_id
        `,
		g.UserID, g.NominationID, g.CategoryID)
	return err
}

func (r *PgxRepository) getGuessByNominationAndUserID(ctx context.Context, userID, nominationID int) (domain.Guess, error) {
	var guess GuessDB
	err := r.db.QueryRow(ctx, `
        SELECT id, user_id, nomination_id
        FROM guesses WHERE nomination_id = $1 AND user_id = $2
        `, nominationID, userID).Scan(&guess.ID, &guess.UserID, &guess.NominationID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Guess{}, pgx.ErrNoRows
		}
		return domain.Guess{}, err
	}

	return toDomainGuess(guess), nil
}

func (r *PgxRepository) fetchCategoryByID(ctx context.Context, categoryID int) (domain.Category, error) {
	var category CategoryDB
	err := r.db.QueryRow(ctx, "SELECT id, name FROM categories WHERE id = $1", categoryID).Scan(&category.ID, &category.Name)

	if err != nil {
		return domain.Category{}, err
	}

	return toDomainCategory(category), nil
}

func (r *PgxRepository) fetchNominationsByCategoryID(ctx context.Context, categoryID int) ([]NominationDB, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, category_id, movie_id, additional_info 
		FROM nominations WHERE category_id = $1`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nominations []NominationDB
	for rows.Next() {
		var n NominationDB
		if err := rows.Scan(&n.ID, &n.CategoryID, &n.MovieID, &n.AdditionalInfo); err != nil {
			return nil, err
		}
		nominations = append(nominations, n)
	}

	return nominations, nil
}

func (r *PgxRepository) fetchMovieByID(ctx context.Context, movieID int) (domain.Movie, error) {
	var movie MovieDB
	err := r.db.QueryRow(ctx, "SELECT id, title, country FROM movies WHERE id = $1", movieID).Scan(&movie.ID, &movie.Title, &movie.Country)
	if err != nil {
		return domain.Movie{}, err
	}
	return toDomainMovie(movie), nil
}

func (r *PgxRepository) fetchPeopleByNominationID(ctx context.Context, nominationID int) ([]domain.Person, error) {
	rows, err := r.db.Query(ctx, `
		SELECT p.id, p.name 
		FROM people p
		JOIN nominee_people np ON p.id = np.person_id
		WHERE np.nomination_id = $1`, nominationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []domain.Person
	for rows.Next() {
		var person PersonDB
		if err := rows.Scan(&person.ID, &person.Name); err != nil {
			return nil, err
		}
		people = append(people, toDomainPerson(person))
	}

	return people, nil
}

func (r *PgxRepository) buildNomination(ctx context.Context, n NominationDB, category domain.Category) (domain.Nomination, error) {
	movie, err := r.fetchMovieByID(ctx, n.MovieID)
	if err != nil {
		return domain.Nomination{}, err
	}

	people, err := r.fetchPeopleByNominationID(ctx, n.ID)
	if err != nil {
		return domain.Nomination{}, err
	}

	return toDomainNomination(n, category, movie, people), nil
}
