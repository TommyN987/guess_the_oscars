package domain

import "time"

type Movie struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Country string `json:"country,omitempty"`
}

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Nomination struct {
	ID             int      `json:"id"`
	Category       Category `json:"category"`
	Movie          Movie    `json:"movie"`
	People         []Person `json:"people"`
	AdditionalInfo string   `json:"additional_info,omitempty"`
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	EmailConfirmed bool   `json:"email_confirmed"`
}

type UserValidation struct {
	ConfirmationToken   string
	ConfirmationExpires time.Time
}

type Guess struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	Nomination int `json:"nomination_id"`
}
