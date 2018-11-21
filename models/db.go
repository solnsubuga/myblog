package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	//Db mongodb collection
	Db *mgo.Database
)

//PostRepository a repository of articles
type PostRepository struct {
	Posts *mgo.Collection
}

//Save creates a new post
func (r *PostRepository) Save(post *Post) error {
	post.ID = bson.NewObjectId()
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now
	err := r.Posts.Insert(post)
	return err
}

//Delete removes a post
func (r *PostRepository) Delete(postID string) error {
	return nil
}

//GetPosts retrieves all posts
func (r *PostRepository) GetPosts() []Post {
	posts := []Post{}
	//TODO: return the error as well
	r.Posts.Find(nil).All(&posts)
	return posts
}

//GetPost retrieves a post by id
func (r *PostRepository) GetPost(id string) (*Post, error) {
	post := Post{}
	err := r.Posts.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&post)
	return &post, err
}
