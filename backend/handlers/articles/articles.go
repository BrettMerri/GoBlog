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

	err := db.C("articles").FindId(bson.ObjectIdHex(id)).One(&article)

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

	err := db.C("articles").Find(nil).Sort("-$natural").All(&articles)

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
	var err error
	if err = c.BindJSON(&json); err == nil {
		db := c.MustGet("db").(*mgo.Database)
		user := models.User{}
		err := db.C("users").FindId(json.User.ID).One(&user)
		if err != nil {
			fmt.Printf("Can't find user, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			article := models.Article{Title: json.Title, Body: json.Body, User: user}
			err := db.C("articles").Insert(&article)
			if err != nil {
				fmt.Printf("Can't add article, go error %v\n", err)
				c.JSON(400, gin.H{
					"result": false,
				})
			} else {
				c.JSON(200, gin.H{
					"result":  true,
					"article": article,
				})
			}
		}
	} else {
		fmt.Printf(err.Error())
		c.JSON(400, err.Error())
	}
}

// Delete : delete article from DB
func Delete(c *gin.Context) {
	var json models.Article
	var err error
	if err = c.BindJSON(&json); err == nil {
		db := c.MustGet("db").(*mgo.Database)
		err := db.C("articles").RemoveId(json.ID)
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
		fmt.Printf(err.Error())
		c.JSON(400, err.Error())
	}
}
