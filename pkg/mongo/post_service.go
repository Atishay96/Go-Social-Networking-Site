package mongo

import (
	mgo "gopkg.in/mgo.v2"

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
