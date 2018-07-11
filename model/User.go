package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	UserName  string        `bson:"user_name" json:"user_name"`
	FirstName string        `bson:"first_name" json:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name"`
	Age       int           `bson:"age" json:"age"`
	Address   string        `bson:"address" json:"address"`
}
