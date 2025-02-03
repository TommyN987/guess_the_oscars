package repository

import "time"

type MovieDB struct {
	ID      int
	Title   string
	Country *string
}

type PersonDB struct {
	ID   int
	Name string
}

type CategoryDB struct {
	ID   int
	Name string
}

type NominationDB struct {
	ID             int
	CategoryID     int
	MovieID        int
	AdditionalInfo *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NomineePersonDB struct {
	ID           int
	NominationID int
	PersonID     int
}

type UserDB struct {
	ID                  int
	Name                string
	Email               string
	EmailConfirmed      bool
	ConfirmationToken   string
	ConfirmationExpires time.Time
	PasswordHash        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type GuessDB struct {
	ID           int
	UserID       int
	NominationID int
	CreatedAt    time.Time
}
