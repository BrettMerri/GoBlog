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
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		nextCommentID := getNextCommentID(db)
		fmt.Println(nextCommentID)
		col := bootstrap(db)

		pushComment := bson.M{"$push": bson.M{"comments": bson.M{"id": nextCommentID, "body": json.Body, "user": json.User}}}

		err := col.Update(bson.M{"_id": json.Article}, pushComment)

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
	} else {
		fmt.Println("Error: Can't add comment")
		c.JSON(400, "Error: Can't add comment")
	}
}

// Delete : delete comment from DB
func Delete(c *gin.Context) {
	var json models.Comment
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		col := bootstrap(db)
		err := col.RemoveId(json.ID)
		if err != nil {
			fmt.Printf("Can't delete comment, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			c.JSON(200, gin.H{
				"result": true,
			})
		}
	} else {
		fmt.Println("Error: Can't delete comment")
		c.JSON(400, "Error: Can't delete comment")
	}
}

func getNextCommentID(db *mgo.Database) int {
	var result bson.M
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counterValue": 1}},
		ReturnNew: true,
	}
	_, err := db.C("counter").Find(bson.M{"_id": "comments"}).Apply(change, &result)

	if err != nil {
		fmt.Printf("Can't get next comment ID , go err %v\n", err)
	}
	newID := result["counterValue"].(int)
	return newID
}

func bootstrap(db *mgo.Database) *mgo.Collection {
	c := db.C("articles")
	return c
}
