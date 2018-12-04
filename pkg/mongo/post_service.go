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

func (ps *PostService) Post(p *root.Post, us root.PostHelper) (root.Post, error) {
	model := postModel{
		ID:        bson.NewObjectId(),
		Text:      p.Text,
		OwnerID:   bson.ObjectIdHex(p.OwnerID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err1 := us.GetUserByID(model.OwnerID.Hex())
	if err1 != nil {
		fmt.Println("ERR 1")
		fmt.Println(err1)
		return root.Post{}, err1
	}
	err := ps.collection.Insert(model)
	if err != nil {
		fmt.Println(err, "err")
		return root.Post{}, err
	}
	return root.Post{
		ID:        model.ID.Hex(),
		OwnerID:   model.OwnerID.Hex(),
		Text:      model.Text,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Owner:     user,
	}, nil
}

func (ps *PostService) GetPosts(limit int, IDs []bson.ObjectId) ([]root.Post, error) {
	var results []postModel
	final := []root.Post{}
	condition := bson.M{
		"_id": bson.M{
			"$nin": IDs,
		},
	}
	err := ps.collection.Find(condition).Sort("-updatedAt").Limit(limit).All(&results)
	if err != nil {
		fmt.Println(err, "err")
		return final, err
	}
	if len(results) == 0 {
		return []root.Post{}, nil
	}
	fmt.Println("results", results)
	for _, v := range results {
		// fmt.Println("results", results)

		final = append(final, root.Post{
			ID:        v.ID.Hex(),
			OwnerID:   v.OwnerID.Hex(),
			Text:      v.Text,
			Comments:  v.Comments,
			Likes:     v.Likes,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return final, nil
}

func (ps *PostService) AddComment(postId string, comment root.Comments) interface{} {
	model := postModel{}
	condition := bson.M{
		"_id": bson.ObjectIdHex(postId),
	}
	err := ps.collection.Find(condition).One(&model)
	if err != nil || model.ID == "" {
		return nil
	}
	change := bson.M{"$push": bson.M{"comments": comment}, "$set": bson.M{"updatedat": time.Now()}}
	err1 := ps.collection.Update(condition, change)
	if err1 != nil {
		fmt.Println("Error ->", err1)
		return nil
	}
	err2 := ps.collection.Find(condition).One(&model)
	if err2 != nil || model.ID == "" {
		return nil
	}
	return model
}

func (ps *PostService) UpdateLike(postId string, like root.Likes) interface{} {
	model := postModel{}
	condition := bson.M{
		"_id": bson.ObjectIdHex(postId),
	}
	err := ps.collection.Find(condition).One(&model)
	if err != nil || model.ID == "" {
		return nil
	}
	flag := false
	var tempID bson.ObjectId
	for _, l := range model.Likes {
		if l.UserID.Hex() == like.UserID.Hex() {
			flag = true
			tempID = l.ID
			break
		}
	}
	fmt.Println("Flag", flag)
	fmt.Println("tempID", tempID)
	var change bson.M
	if flag == false {
		change = bson.M{"$push": bson.M{"likes": like}, "$set": bson.M{"updatedat": time.Now()}}
	} else {
		change = bson.M{"$pull": bson.M{"likes": bson.M{"_id": tempID}}, "$set": bson.M{"updatedat": time.Now()}}
	}
	err1 := ps.collection.Update(condition, change)
	if err1 != nil {
		fmt.Println("Error ->", err1)
		return nil
	}
	err2 := ps.collection.Find(condition).One(&model)
	if err2 != nil || model.ID == "" {
		return nil
	}
	return model
}
