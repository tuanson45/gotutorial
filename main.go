package main

import (
	"github.com/kataras/iris"
	"devlife.info/gotutorial/router"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	app := iris.New()

	/**
		Logger config
	 */
	customLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Columns:            true,
		MessageContextKeys: []string{"logger_message"},
	})
	app.Use(customLogger)

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

	app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		ctx.Values().Set("logger_message",
			"a dynamic message passed to the logs")
		ctx.Writef("My Custom error page %s", errMessage)
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
