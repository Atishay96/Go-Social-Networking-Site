package root

import "time"

type User struct {
	ID                   string
	Username             string
	Password             string
	FirstName            string
	LastName             string
	Email                string
	PhoneNumber          string
	PhoneNumberExtension string
	DOB                  time.Time
	AboutMe              string
	Verified             bool
	Blocked              bool
	VerificationSecret   string
	UpdatedAt            time.Time
	CreatedAt            time.Time
	LastLoggedIn         time.Time
}

type UserService interface {
	CreateUser(u *User) error
	CheckUserName(username string) bool
	CheckEmail(email string) bool
	HandleSecret(secret string) (User, error)
	UpdateUser(fields []string, VerificationSecret string, email string) error
	GetUserByUsername(username string) (User, error)
	GetUserByID(ID string) (User, error)
	CheckStatus(email string) (bool, User)
	GetUserByParams(params []string) interface{}
	GetOtherUserByParams(ID string) interface{}
	CheckToken(params []string) bool
	UpdateLastLoggedIn(ID string) (User, error)
}
