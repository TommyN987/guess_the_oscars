package api

import (
	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
)

type NominationResponse struct {
	ID             int             `json:"id"`
	Movie          domain.Movie    `json:"movie"`
	People         []domain.Person `json:"people"`
	AdditionalInfo string          `json:"additional_info,omitempty"`
	IsGuessed      bool            `json:"is_guessed"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GuessDTO struct {
	NominationID int `json:"nomination_id"`
	CategoryID   int `json:"category_id"`
}
