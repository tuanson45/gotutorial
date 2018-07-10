package router

import (
	"github.com/kataras/iris"
	"devlife.info/gotutorial/controller"
)

func ApiRouter(app *iris.Application) {
	app.Get("/user", controller.All)
	app.Get("/user/{user_id:string}", controller.Get)
	app.Post("/user", controller.Post)
}
