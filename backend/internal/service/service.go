package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
	"github.com/TommyN987/guess_the_oscars/backend/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Service interface {
	CheckHealth(ctx context.Context) (string, error)
	RegisterUser(ctx context.Context, user domain.User) error
	ValidateToken(ctx context.Context, token string) (domain.User, error)
	LoginUser(ctx context.Context, email, password string) (string, domain.User, error)
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetNominationsByCategory(ctx context.Context, userID, categoryID int) ([]domain.Nomination, *domain.Guess, error)
	SubmitGuess(ctx context.Context, guess domain.Guess) error
}

type DefaultService struct {
	repo repository.Repository
}

func NewDefaultService(repo repository.Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

func (s *DefaultService) CheckHealth(ctx context.Context) (string, error) {
	return s.repo.CheckHealth(ctx)
}

func (s *DefaultService) RegisterUser(ctx context.Context, user domain.User) error {
	// Check if the user already exists
	_, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		// No error means the user exists
		return errors.New("User already exists.")
	} else if err != pgx.ErrNoRows {
		// An unexpected error occurred (e.g., database issue)
		return err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	token, err := generateValidationToken()
	if err != nil {
		return err
	}

	userValidation := domain.UserValidation{
		ConfirmationToken:   token,
		ConfirmationExpires: time.Now().Add(24 * time.Hour),
	}

	err = s.repo.CreateUser(ctx, user, userValidation)
	if err != nil {
		return err
	}

	go func() {
		if err := s.sendValidationEmail(user.Name, user.Email, token); err != nil {
			log.Printf("Failed to send validation email: %v", err) // Ensure errors are logged
		} else {
			log.Println("Validation email sent successfully.") // Confirmation log
		}
	}()

	return nil
}

func (s *DefaultService) ValidateToken(ctx context.Context, token string) (domain.User, error) {
	user, err := s.repo.GetUserByToken(ctx, token)
	if err != nil {
		return user, fmt.Errorf("Invalid token: %s", err)
	}

	err = s.repo.UpdateUserConfirmation(ctx, user.ID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *DefaultService) LoginUser(ctx context.Context, email, password string) (string, domain.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", domain.User{}, errors.New("Invalid email or password...!")
	}

	if !user.EmailConfirmed {
		return "", domain.User{}, errors.New("Please confirm your email before logging in.")
	}

	if err := checkPasswordHash(password, user.Password); err != nil {
		return "", domain.User{}, errors.New("Invalid email or password.")
	}

	token, e := GenerateJWT(user)
	if e != nil {
		return "", domain.User{}, e
	}

	return token, user, nil
}

func (s *DefaultService) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *DefaultService) GetNominationsByCategory(ctx context.Context, userID, categoryID int) ([]domain.Nomination, *domain.Guess, error) {
	return s.repo.GetNominationsByCategoryID(ctx, userID, categoryID)
}

func (s *DefaultService) SubmitGuess(ctx context.Context, guess domain.Guess) error {
	return s.repo.UpsertGuess(ctx, guess)
}

func (s *DefaultService) sendValidationEmail(name, email, token string) error {
	validationLink := fmt.Sprintf("http://localhost:5173/validate?token=%s", token)
	emailBody := fmt.Sprintf("Click the link to confirm your email: %s", validationLink)

	toEmail := mail.Email{
		Name:    name,
		Address: email,
	}

	return SendEmail(&toEmail, "Validation", emailBody)

}

func generateValidationToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
