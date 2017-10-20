package comments

import (
	"fmt"

	"github.com/brettmerri/GoBlog/backend/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Add : add comment to DB
func Add(c *gin.Context) {
	var json models.Comment
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
			comment := models.Comment{ID: bson.NewObjectId(), Body: json.Body, User: user}
			pushComment := bson.M{"$push": bson.M{"comments": comment}}

			err := db.C("articles").Update(bson.M{"_id": json.Article}, pushComment)

			if err != nil {
				fmt.Printf("Can't add comment, go error %v\n", err)
				c.JSON(400, gin.H{
					"result": false,
					"err":    err,
				})
			} else {
				c.JSON(200, gin.H{
					"result":  true,
					"comment": comment,
				})
			}
		}
	} else {
		fmt.Println(err.Error())
		c.JSON(400, err.Error())
	}
}

// Delete : delete comment from DB
func Delete(c *gin.Context) {
	var json models.Comment
	var err error
	if err = c.BindJSON(&json); err == nil {
		db := c.MustGet("db").(*mgo.Database)

		pullComment := bson.M{"$pull": bson.M{"comments": bson.M{"_id": json.ID}}}

		err := db.C("articles").Update(bson.M{"_id": json.Article}, pullComment)

		if err != nil {
			fmt.Printf("Can't remove comment, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
				"err":    err,
			})
		} else {
			c.JSON(200, gin.H{
				"result": true,
			})
		}
	} else {
		fmt.Println(err.Error())
		c.JSON(400, err.Error())
	}
}
