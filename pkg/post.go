package root

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comments struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  bson.ObjectId `bson:"UserId,omitempty"`
	Comment string
}

type Likes struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserID bson.ObjectId `bson:"UserId,omitempty"`
}

type PostHelper interface {
	GetUserByID(ID string) (User, error)
}

type Post struct {
	ID        string
	OwnerID   string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comments
	Likes     []Likes
	IDs       []string
	Limit     int
	Comment   string
	Owner     User
}

type PostService interface {
	Post(p *Post, us PostHelper) (Post, error)
	GetPosts(limit int, IDs []bson.ObjectId, us PostHelper) ([]Post, error)
	AddComment(postId string, comment Comments) interface{}
	UpdateLike(postId string, like Likes) interface{}
}
