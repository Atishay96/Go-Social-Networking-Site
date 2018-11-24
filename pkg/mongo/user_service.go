package mongo

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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

func (p *UserService) CheckUserName(username string) bool {
	model := userModel{}
	// use regex
	p.collection.Find(bson.M{"username": username}).One(&model)
	if model.Username != "" {
		return false
	}
	return true
}

func (p *UserService) CheckEmail(email string) bool {
	model := userModel{}
	p.collection.Find(bson.M{"email": strings.ToLower(email)}).One(&model)
	if model.Email != "" {
		return false
	}
	return true
}

func (p *UserService) HandleSecret(secret string) error {
	condition := bson.M{
		"$and": []bson.M{
			bson.M{"verificationsecret": secret},
			bson.M{"verified": false},
		},
	}
	change := bson.M{"$set": bson.M{"verified": true, "verifiedon": time.Now(), "updatedat": time.Now()}}
	err := p.collection.Update(condition, change)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
