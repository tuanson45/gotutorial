package controller

import (
	"github.com/kataras/iris/context"
	"devlife.info/gotutorial/model"
	"gopkg.in/mgo.v2/bson"
	"devlife.info/gotutorial/dao"
	"log"
	"devlife.info/gotutorial/db"
)

func All(context context.Context) {
	doe := model.User{
		UserName:  "Johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Address:   "Neither FBI knows!!!",
		Age:       25,
	}
	context.JSON(doe)
}

func Get(context context.Context) {
	userId := context.Params().Get("user_id")
	context.Text(userId)
}

func PostUser(context context.Context) {
	var user model.User
	database := db.Init()
	userDao := dao.UserDao{DB: database}
	context.ReadJSON(&user)
	user.ID = bson.NewObjectId()
	err := userDao.Insert(user)
	if err != nil {
		log.Fatal(err)
		context.Text("err")
	} else {
		context.Text("Insert success!")
	}
}
