package a

import "time"

type User struct {
	UserID    int       `json:"userId"`    // OK
	FirstName string    `json:"firstName"` // OK
	LastName  string    `json:"-"`         // OK
	Username  string    `json:user-name`   // nocamel OK
	BirthDay  string    `json:"birth_day"` // want "invalid JSON tag birth_day"
	CreatedAt time.Time `json:"_"`         // want "invalid JSON tag _"
}
