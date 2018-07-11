package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	SERVER   = "localhost"
	DATABASE = "golang_database"
)

func Connect() *mgo.Session {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	return session
}
