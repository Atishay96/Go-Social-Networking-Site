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

	model := postModel{
		Text:      p.Text,
		OwnerID:   bson.ObjectIdHex(p.OwnerID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := ps.collection.Insert(model)
	if err != nil {
		fmt.Println(err, "err")
		return root.Post{
			ID:        model.ID.Hex(),
			Text:      model.Text,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}, err
	}
	return root.Post{
		ID:        model.ID.Hex(),
		Text:      model.Text,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}
