package controller

import (
	"github.com/kataras/iris/context"
	"devlife.info/gotutorial/model"
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

func Post(context context.Context) {
	var user model.User
	context.ReadJSON(&user)
	context.JSON(user)
}
