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
	id := c.Param("id")
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}

	col := bootstrap(db)

	err := col.FindId(bson.ObjectIdHex(id)).One(&article)

	if err != nil {
		fmt.Printf("Can't find article, go error %v\n", err)
		c.JSON(400, err)
	} else {
		c.JSON(200, article)
	}
}

// ReadAll : read all articles from DB
func ReadAll(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	articles := []models.Article{}

	col := bootstrap(db)

	err := col.Find(nil).All(&articles)

	if err != nil {
		fmt.Printf("Can't find articles, go error %v\n", err)
		c.JSON(400, err)
	} else {
		c.JSON(200, articles)
	}
}

// Add : add article from DB
func Add(c *gin.Context) {
	var json models.Article
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		col := bootstrap(db)
		err := col.Insert(models.Article{Title: json.Title, Body: json.Body, User: json.User})
		if err != nil {
			fmt.Printf("Can't add article, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			c.JSON(200, gin.H{
				"result": true,
			})
		}
	} else {
		fmt.Printf("Error: Can't add article")
		c.JSON(400, "Error: Can't add article")
	}
}

// Delete : delete article from DB
func Delete(c *gin.Context) {
	var json models.Article
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		col := bootstrap(db)
		err := col.RemoveId(json.ID)
		if err != nil {
			fmt.Printf("Can't delete article, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			c.JSON(200, gin.H{
				"result": true,
			})
		}
	} else {
		fmt.Printf("Error: Can't delete article")
		c.JSON(400, "Error: Can't delete article")
	}
}

func bootstrap(db *mgo.Database) *mgo.Collection {
	c := db.C("articles")

	index := mgo.Index{
		Key:        []string{"title"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
