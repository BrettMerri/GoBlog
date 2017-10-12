package articles

import (
	"fmt"

	"github.com/brettmerri/GoBlog/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Read(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}

	err := db.C(models.CollectionArticle).Find(bson.M{"title": "Ale"}).One(&article)

	if err != nil {
		fmt.Printf("Can't find article, go error %v\n", err)
		panic(err.Error())
	}

	c.JSON(200, article)
}

func Add(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	err := db.C(models.CollectionArticle).Insert(models.Article{Title: "Ale", Body: "+55 53 1234 4321"},
		models.Article{Title: "Cla", Body: "+66 33 1234 5678"})

	if err != nil {
		fmt.Printf("Can't add article, go error %v\n", err)
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"result": true,
	})
}
