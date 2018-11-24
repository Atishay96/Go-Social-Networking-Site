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
}

type UserService interface {
	CreateUser(u *User) error
	CheckUserName(username string) bool
	CheckEmail(email string) bool
	HandleSecret(secret string) error
}
