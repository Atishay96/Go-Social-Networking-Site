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
	email = strings.ToLower(email)
	p.collection.Find(bson.M{"email": strings.ToLower(email)}).One(&model)
	if model.Email != "" {
		return false
	}
	return true
}

func (p *UserService) HandleSecret(secret string) (root.User, error) {
	model := userModel{}
	condition := bson.M{
		"$and": []bson.M{
			bson.M{"verificationsecret": secret},
			bson.M{"verified": false},
		},
	}
	change := bson.M{"$set": bson.M{"verified": true, "verifiedon": time.Now(), "updatedat": time.Now()}}
	err := p.collection.Update(condition, change)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return root.User{
			ID:        model.ID.Hex(),
			Username:  model.Username,
			Password:  "-",
			UpdatedAt: model.UpdatedAt}, err
	}
	err1 := p.collection.Find(bson.M{"verificationsecret": secret}).One(&model)
	if err1 != nil {
		fmt.Println("err1")
		fmt.Println(err1)
		return root.User{
			ID:        model.ID.Hex(),
			Username:  model.Username,
			Password:  "-",
			UpdatedAt: model.UpdatedAt}, err1
	}
	return root.User{
		ID:        model.ID.Hex(),
		Username:  model.Username,
		Password:  "-",
		UpdatedAt: model.UpdatedAt}, nil
}

func (p *UserService) UpdateUser(fields []string, VerificationSecret string, email string) error {
	// make it generic

	email = strings.ToLower(email)

	condition := bson.M{
		"$and": []bson.M{
			bson.M{"email": email},
			bson.M{"verified": false},
		},
	}
	change := bson.M{
		"$set": bson.M{
			"verificationsecret": VerificationSecret,
			"updatedat":          time.Now(),
		},
	}
	err := p.collection.Update(condition, change)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (p *UserService) GetUserByUsername(username string) (root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return root.User{
		ID:       model.ID.Hex(),
		Username: model.Username,
		Password: "-"}, err
}

func (p *UserService) GetUserByID(ID string) (root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"ID": ID}).One(&model)
	return root.User{
		ID:       model.ID.Hex(),
		Username: model.Username,
		Password: "-"}, err
}

func (p *UserService) CheckStatus(email string) (bool, root.User) {
	model := userModel{}
	condition := bson.M{
		"$and": []bson.M{
			bson.M{
				"verified": true,
			},
			bson.M{
				"blocked": false,
			},
		},
	}
	change := bson.M{
		"$set": bson.M{
			"updatedat": time.Now(),
		},
	}
	err1 := p.collection.Update(condition, change)
	if err1 != nil {
		fmt.Println(err1)
		return false, root.User{
			ID:       model.ID.Hex(),
			Username: model.Username,
			Password: "-",
		}
	}
	err := p.collection.Find(condition).One(&model)
	if err != nil || model.Username == "" {
		return false, root.User{
			ID:       model.ID.Hex(),
			Username: model.Username,
			Password: "-",
		}
	}
	return true, root.User{
		ID:        model.ID.Hex(),
		Username:  model.Username,
		Password:  "-",
		UpdatedAt: model.UpdatedAt,
	}
}

func (p *UserService) GetUserByParams(params []string) interface{} {
	model := userModel{}
	// make it dynamic. Loop it
	fmt.Println(params[0])
	fmt.Println(params[1])
	fmt.Println(params[2])
	updatedat, err := time.Parse(time.RFC3339, params[2])
	if err != nil {
		fmt.Println("Error occured", err)
		return nil
	}
	condition := bson.M{
		"$and": []bson.M{
			bson.M{
				"username": params[0],
			},
			bson.M{
				"_id": bson.ObjectIdHex(params[1]),
			},
			bson.M{
				"updatedat": updatedat,
			},
		},
	}
	err1 := p.collection.Find(condition).One(&model)
	if err1 != nil || model.Username == "" {
		return nil
	}
	return root.User{
		ID:                   model.ID.Hex(),
		Username:             model.Username,
		Password:             "-",
		UpdatedAt:            model.UpdatedAt,
		FirstName:            model.FirstName,
		LastName:             model.LastName,
		CreatedAt:            model.CreatedAt,
		Email:                model.Email,
		PhoneNumber:          model.PhoneNumber,
		PhoneNumberExtension: model.PhoneNumberExtension,
		DOB:                  model.DOB,
		AboutMe:              model.AboutMe,
	}
}

func (p *UserService) GetOtherUserByParams(ID string) interface{} {
	model := userModel{}
	// merge it with GetUserByParams
	condition := bson.M{
		"_id": bson.ObjectIdHex(ID),
	}
	err1 := p.collection.Find(condition).One(&model)
	if err1 != nil || model.Username == "" {
		return nil
	}
	return root.User{
		ID:                   model.ID.Hex(),
		Username:             model.Username,
		Password:             "-",
		UpdatedAt:            model.UpdatedAt,
		FirstName:            model.FirstName,
		LastName:             model.LastName,
		CreatedAt:            model.CreatedAt,
		Email:                model.Email,
		PhoneNumber:          model.PhoneNumber,
		PhoneNumberExtension: model.PhoneNumberExtension,
		DOB:                  model.DOB,
		AboutMe:              model.AboutMe,
	}
}
