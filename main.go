package main

import (
	"github.com/kataras/iris"
	"devlife.info/gotutorial/router"
)

func main() {
	app := iris.New()

	/**
		Handle error
	 */
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})

	/**
		Log router request
	 */
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	router.ApiRouter(app)

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}
