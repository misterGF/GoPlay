package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	// Get method - pass string back
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Welcome to the iris!")
	})

	// Shortcut with get method. Write string back
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("PONG!")
	})

	// Get with a json response
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hey buddy!"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
