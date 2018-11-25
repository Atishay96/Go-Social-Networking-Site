package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type postModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	OwnerID   bson.ObjectId `bson:"ownerId,omitempty"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []string //slices of slices
	Likes     []string //slices of slices
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
