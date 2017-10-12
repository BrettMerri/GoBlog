package articles

import (
	"fmt"

	"github.com/brettmerri/GoBlog/backend/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Read : read article from DB
func Read(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}

	col := bootstrap(db)

	err := col.Find(bson.M{"title": "Ale"}).One(&article)

	if err != nil {
		fmt.Printf("Can't find article, go error %v\n", err)
		panic(err.Error())
	}

	c.JSON(200, article)
}

// Add : add article from DB
func Add(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	col := bootstrap(db)

	err := col.Insert(models.Article{Title: "Ale", Body: "+55 53 1234 4321"},
		models.Article{Title: "Cla", Body: "+66 33 1234 5678"})

	if err != nil {
		fmt.Printf("Can't add article, go error %v\n", err)
		c.JSON(200, err)
	}

	c.JSON(200, gin.H{
		"result": true,
	})
}

func bootstrap(db *mgo.Database) *mgo.Collection {
	c := db.C(models.CollectionArticle)

	index := mgo.Index{
		Key:        []string{"title"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
