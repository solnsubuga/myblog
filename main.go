package main

import (
	"github.com/gin-gonic/gin"
	"github.com/solnsubuga/myblog/handlers"
	"github.com/solnsubuga/myblog/models"
	"gopkg.in/mgo.v2"
)

const (
	mongoURL        = "mongodb://localhost:27017"
	dbName          = "myblog"
	postsCollection = "posts"
)

func main() {
	//TODO: use viper to read from configs
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	defer session.Close()
	models.Db = session.DB(dbName)
	postsCol := models.Db.C(postsCollection)

	postRepo := models.PostRepository{Posts: postsCol}
	postsHandler := handlers.PostHandler{Repo: &postRepo}

	//TODO abstract this way in a function
	r := gin.Default()

	//handle request to api v1
	api := r.Group("/api/v1")
	{
		api.GET("/posts", postsHandler.HandlePosts)
		api.POST("/posts", postsHandler.CreatePost)
		api.GET("/post/:id", postsHandler.GetPost)
	}

	r.Run(":8000")

}
