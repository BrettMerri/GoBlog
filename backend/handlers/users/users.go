package users

import (
	"fmt"

	"github.com/brettmerri/GoBlog/backend/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Read : read user from DB
func Read(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}

	col := bootstrap(db)

	err := col.FindId(bson.ObjectIdHex(id)).One(&user)

	if err != nil {
		fmt.Printf("Can't find user, go error %v\n", err)
		c.JSON(400, err)
	} else {
		c.JSON(200, user)
	}
}

// ReadAll : read all users from DB
func ReadAll(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	users := []models.User{}

	col := bootstrap(db)

	err := col.Find(nil).All(&users)

	if err != nil {
		fmt.Printf("Can't find users, go error %v\n", err)
		c.JSON(400, err)
	} else {
		c.JSON(200, users)
	}
}

// Add : add user from DB
func Add(c *gin.Context) {
	var json models.User
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		col := bootstrap(db)
		err := col.Insert(models.User{Username: json.Username})
		if err != nil {
			fmt.Printf("Can't add user, go error %v\n", err)
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

// Delete : delete user from DB
func Delete(c *gin.Context) {
	var json models.User
	if c.BindJSON(&json) == nil {
		db := c.MustGet("db").(*mgo.Database)
		col := bootstrap(db)
		err := col.RemoveId(json.ID)
		if err != nil {
			fmt.Printf("Can't delete user, go error %v\n", err)
			c.JSON(400, gin.H{
				"result": false,
			})
		} else {
			c.JSON(200, gin.H{
				"result": true,
			})
		}
	} else {
		fmt.Printf("Error: Can't delete user")
		c.JSON(400, "Error: Can't delete user")
	}
}

func bootstrap(db *mgo.Database) *mgo.Collection {
	c := db.C("users")

	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
