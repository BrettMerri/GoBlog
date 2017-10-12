package articles

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"github.com/brettmerri/GoBlog/models"
)

func read(dbCollection *mgo.Collection) {

	article := models.Article{}

	err := dbCollection.Find(bson.M{"title": "Ale"}).One(&result)

	if err != nil {
		fmt.Printf("Can't find article, go error %v\n", err)
		panic(err.Error())
	}

	c.JSON(200, article)

	return result
}

func add(dbCollection *mgo.Collection) {

	article := models.Article{}

	err := dbCollection.Insert(&article{Title: "Ale", Body: "+55 53 1234 4321"},
		&article{Title: "Cla", Body: "+66 33 1234 5678"})

	if err != nil {
		fmt.Printf("Can't add article, go error %v\n", err)
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"result": true
	})
}
