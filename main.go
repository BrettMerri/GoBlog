package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/brettmerri/goblog/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name      string
	Phone     string
	Timestamp time.Time
}

var database = "go-development"

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

	dbCollection := bootstrap(session)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		b := read(dbCollection)
		c.JSON(200, b)
	})

	router.GET("/add", func(c *gin.Context) {
		b := add(dbCollection)
		c.JSON(200, gin.H{
			"result": b,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
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

func read(dbCollection *mgo.Collection) Person {

	result := Person{}

	err := dbCollection.Find(bson.M{"name": "Ale"}).One(&result)

	if err != nil {
		panic(err)
	}

	return result
}

func add(dbCollection *mgo.Collection) bool {

	err := dbCollection.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	return true
}
