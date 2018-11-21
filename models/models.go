package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Author creator of articles
type Author struct {
	ID        bson.ObjectId `json:"id"`
	FirstName string        `json:"FirstName"`
	LastName  string        `json:"LastName"`
	Password  string        `json:"password"`
	Username  string        `json:"username"`
}

//Post article
type Post struct {
	ID        bson.ObjectId `json:"id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	AuthorID  bson.ObjectId `json:"authorId,omitempty"`
}
