package mongo

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"Go-Social/pkg"
)

type PostService struct {
	collection *mgo.Collection
}

func NewPostService(session *mgo.Session, config *root.MongoConfig) *PostService {
	collection := session.DB(config.DbName).C("posts")
	collection.EnsureIndex(postModelIndex())
	return &PostService{collection}
}

func (ps *PostService) Post(p *root.Post) (root.Post, error) {

	model := root.Post{
		Text:      p.Text,
		OwnerID:   p.OwnerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newModel := postModel{}
	err := ps.collection.Insert(model)
	if err != nil {
		fmt.Println(err, "err")
		return model, err
	}
	condition := bson.M{}
	err1 := ps.collection.Find(condition).One(&newModel)
	if err1 != nil {
		fmt.Println(err1)
		return root.Post{
			ID:        newModel.ID.Hex(),
			Text:      newModel.Text,
			CreatedAt: newModel.CreatedAt,
			UpdatedAt: newModel.UpdatedAt,
		}, err1
	}
	// send the correct ID
	return root.Post{
		ID:        newModel.ID.Hex(),
		Text:      newModel.Text,
		CreatedAt: newModel.CreatedAt,
		UpdatedAt: newModel.UpdatedAt,
	}, nil
}
