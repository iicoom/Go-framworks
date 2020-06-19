package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })

    app.Listen(":8080")
}