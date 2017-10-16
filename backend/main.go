package main

import (
	"github.com/BrettMerri/GoBlog/backend/handlers/users"
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
		article := api.Group("/article")
		{
			article.GET("/read", articles.ReadAll)
			article.GET("/read/:id", articles.Read)
			article.POST("/add", articles.Add)
			article.POST("/delete", articles.Delete)
		}
		user := api.Group("/user")
		{
			user.GET("/read", users.ReadAll)
			user.GET("/read/:id", users.Read)
			user.POST("/add", users.Add)
			user.POST("/delete", users.Delete)
		}
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
