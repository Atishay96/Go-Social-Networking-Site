package mongo

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"Go-Social/pkg"
)

type userModel struct {
	ID                   bson.ObjectId `bson:"_id,omitempty"`
	Username             string
	PasswordHash         string
	Salt                 string
	FirstName            string
	LastName             string
	Email                string
	PhoneNumber          string
	PhoneNumberExtension string
	DOB                  time.Time
	AboutMe              string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Verified             bool
	VerifiedOn           time.Time
	Blocked              bool
	BlockedOn            time.Time
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username", "id", "email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *root.User) (*userModel, error) {
	user := userModel{Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email}
	err := user.setSaltedPassword(u.Password)
	return &user, err
}

func (u *userModel) comparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *userModel) setSaltedPassword(password string) error {
	salt := uuid.New().String()
	passwordBytes := []byte(password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash[:])
	u.Salt = salt

	return nil
}
