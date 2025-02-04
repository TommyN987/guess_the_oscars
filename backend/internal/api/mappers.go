package api

import "github.com/TommyN987/guess_the_oscars/backend/internal/domain"

func toResponseNomination(n domain.Nomination) NominationResponse {
	return NominationResponse{
		ID:             n.ID,
		Movie:          n.Movie,
		People:         n.People,
		AdditionalInfo: n.AdditionalInfo,
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
