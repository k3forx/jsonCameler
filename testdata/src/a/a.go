package a

import "time"

type User struct {
	UserID    int       `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"-"`
	Username  string    `json:"user-name"` // nocamel
	BirthDay  string    `json:"birth_day"` // want "invalid JSON tag `birth_day`"
	CreatedAt time.Time `json:"_"`         // want "invalid JSON tag `_`"
}

type UserAddress struct {
	ID        int
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at" some tags` // want "invalid JSON tag `created_at`"
}

type Config struct {
	APIKey string `env:"API_KEY"`
}
