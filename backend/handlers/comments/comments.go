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
		nextCommentID := getNextCommentID(db)
		user := models.User{}
		err := db.C("users").FindId(json.User.ID).One(&user)
		if err != nil {
			fmt.Printf("Can't find user, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			pushComment := bson.M{"$push": bson.M{"comments": bson.M{"id": nextCommentID, "body": json.Body, "user": user}}}

			err := db.C("articles").Update(bson.M{"_id": json.Article}, pushComment)

			if err != nil {
				fmt.Printf("Can't add comment, go error %v\n", err)
				c.JSON(400, gin.H{
					"result": false,
				})
			} else {
				c.JSON(200, gin.H{
					"result": true,
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

		pullComment := bson.M{"$pull": bson.M{"comments": bson.M{"id": json.ID}}}

		err := db.C("articles").Update(bson.M{"_id": json.Article}, pullComment)

		if err != nil {
			fmt.Printf("Can't remove comment, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
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

func getNextCommentID(db *mgo.Database) int {
	var result bson.M
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counterValue": 1}},
		ReturnNew: true,
	}
	_, err := db.C("counter").Find(bson.M{"_id": "comments"}).Apply(change, &result)

	if err != nil { // If counter collection with id comments doesnt exist, create it
		db.C("counter").Insert(bson.M{"_id": "comments", "counterValue": 1})
		return 1
	}
	newID := result["counterValue"].(int)
	return newID
}
