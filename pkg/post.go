package root

import "gopkg.in/mgo.v2/bson"

type Post struct {
	ID          bson.ObjectId
	ownerID     bson.ObjectId
	attachments []string
	text        string
	createdAt   string
	updatedAt   string
	comments    []string
	likes       []string
}

type PostService interface {
}
