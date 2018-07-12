package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"devlife.info/gotutorial/model"
)

type UserDao struct {
	DB      *mgo.Database
}

const (
	COLLECTION = "users"
)


func (m *UserDao) FindAll() ([]model.User, error) {
	var user []model.User
	err := m.DB.C(COLLECTION).Find(bson.M{}).All(&user)
	return user, err
}

func (m *UserDao) FindById(id string) (model.User, error) {
	var user model.User
	err := m.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserDao) Insert(user model.User) error {
	err := m.DB.C(COLLECTION).Insert(&user)
	return err
}

func (m *UserDao) Delete(user model.User) error {
	err := m.DB.C(COLLECTION).Remove(&user)
	return err
}

func (m *UserDao) Update(user model.User) error {
	err := m.DB.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
