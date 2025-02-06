package api

import "github.com/TommyN987/guess_the_oscars/backend/internal/domain"

func toResponseNomination(n domain.Nomination, guess *domain.Guess) NominationResponse {
	isGuessed := false
	if guess != nil && guess.NominationID == n.ID {
		isGuessed = true
	}

	return NominationResponse{
		ID:             n.ID,
		Movie:          n.Movie,
		People:         n.People,
		AdditionalInfo: n.AdditionalInfo,
		IsGuessed:      isGuessed,
	}
}

func toDomainUser(u UserDTO) domain.User {
	return domain.User{
		Name:           u.Name,
		Email:          u.Email,
		Password:       u.Password,
		EmailConfirmed: false,
	}
}

func toDomainGuess(g GuessDTO, userID int) domain.Guess {
	return domain.Guess{
		UserID:       userID,
		NominationID: g.NominationID,
		CategoryID:   g.CategoryID,
	}
}
