package dao

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"devlife.info/gotutorial/model"
)

type UserDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UserDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *UserDao) FindAll() ([]model.User, error) {
	var user []model.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&user)
	return user, err
}

func (m *UserDao) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserDao) Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *UserDao) Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *UserDao) Update(user model.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
