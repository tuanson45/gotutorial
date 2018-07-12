package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	SERVER   = "localhost:27017"
	DATABASE = "golang_database"
)

func Init() (*mgo.Database) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(DATABASE)
}
