package repository

import "github.com/TommyN987/guess_the_oscars/backend/internal/domain"

func toDomainCategory(c CategoryDB) domain.Category {
	return domain.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

func toDomainMovie(m MovieDB) domain.Movie {
	return domain.Movie{
		ID:      m.ID,
		Title:   m.Title,
		Country: getString(m.Country),
	}
}

func toDomainPerson(p PersonDB) domain.Person {
	return domain.Person{
		ID:   p.ID,
		Name: p.Name,
	}
}

func toDomainNomination(n NominationDB, c domain.Category, m domain.Movie, p []domain.Person) domain.Nomination {
	return domain.Nomination{
		ID:             n.ID,
		Category:       c,
		Movie:          m,
		People:         p,
		AdditionalInfo: getString(n.AdditionalInfo),
	}
}

func toDomainGuess(g GuessDB) domain.Guess {
	return domain.Guess{
		NominationID: g.NominationID,
		UserID:       g.UserID,
	}
}

func getString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
