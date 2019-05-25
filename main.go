package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"log"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/xxxxssss12/mvc_exp/common"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	app.Any("/", route)
	// Method:   GET
	// Resource: http://localhost:8080
	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	ctx.HTML("<h1>Welcome</h1>")
	//})
	//
	//// same as app.Handle("GET", "/ping", [...])
	//// Method:   GET
	//// Resource: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("pong")
	//})
	//
	//// Method:   GET
	//// Resource: http://localhost:8080/hello
	//app.Get("/", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})s

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func route(ctx iris.Context) {
	var result = common.BuildSuccHttpResult(iris.Map{"message": "Hello Iris!", "url": ctx.Request().URL.String()})
	jsonStr, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result->" + string(jsonStr))
	ctx.JSON(result)
}
