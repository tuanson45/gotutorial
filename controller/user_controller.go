package controller

import (
	"github.com/kataras/iris/context"
	"devlife.info/gotutorial/model"
)

func Get(context context.Context) {
	doe := model.User{
		UserName:  "Johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Address:   "Neither FBI knows!!!",
		Age:       25,
	}

	context.JSON(doe)
}

func Post(context context.Context) {
	var user model.User
	context.ReadJSON(&user)
	context.JSON(user)
}
