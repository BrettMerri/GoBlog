package main

import (
	"github.com/brettmerri/GoBlog/db"
	"github.com/brettmerri/GoBlog/handlers/articles"
	"github.com/brettmerri/GoBlog/middlewares"
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
	router.Use(middlewares.ErrorHandler)

	router.GET("/article", articles.Read)
	router.GET("/add", articles.Add)

	router.Run() // listen and serve on 0.0.0.0:8080
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
