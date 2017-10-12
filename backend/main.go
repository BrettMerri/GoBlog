package main

import (
	"github.com/brettmerri/GoBlog/backend/db"
	"github.com/brettmerri/GoBlog/backend/handlers/articles"
	"github.com/brettmerri/GoBlog/backend/middlewares"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var database = "go-development"

func init() {
	db.Connect()
}

func main() {

	router := gin.Default()

	router.Use(middlewares.Connect)

	router.Use(static.Serve("/", static.LocalFile("./public", false)))

	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("/article", articles.Read)
		api.GET("/add", articles.Add)
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func bootstrap(s *mgo.Session) *mgo.Collection {

	c := s.DB(database).C("people")
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
