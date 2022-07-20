package a

type User struct {
	UserID    int    `json:"userId"`    // OK
	FirstName string `json:"firstName"` // OK
	LastName  string `json:"-"`         // OK
	BirthDay  string `json:"birth_day"` // want "invalid JSON tag birth_day"
}
