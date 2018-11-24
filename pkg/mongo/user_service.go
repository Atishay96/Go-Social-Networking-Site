package mongo

import (
	"gopkg.in/mgo.v2"

	"Go-Social/pkg"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mgo.Session, config *root.MongoConfig) *UserService {
	collection := session.DB(config.DbName).C("user")
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (p *UserService) CreateUser(u *root.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}
	return p.collection.Insert(&user)
}
