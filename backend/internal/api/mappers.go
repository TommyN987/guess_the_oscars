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
