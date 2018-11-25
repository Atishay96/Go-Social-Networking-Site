package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type postModel struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	ownerID bson.ObjectId `bson:"ownerId,omitempty"`
	// location
	attachments []string
	text        string
	createdAt   string
	updatedAt   string
	comments    []string //slices of slices
	likes       []string //slices of slices
}

func postModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}
