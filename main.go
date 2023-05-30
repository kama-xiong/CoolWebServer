package main

import (
	"WebServer/web/handlers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	//设置模板
	app.RegisterView(iris.Django("./web/views", ".html").Reload(true))
	//加载静态文件
	app.HandleDir("/static", "./web/static")
	app.PartyFunc("/product", handlers.ProductPartyHandler)
	app.Get("/about", handlers.AboutHandler)
	app.Get("/", func(ctx *context.Context) {
		ctx.Redirect("/about")
	})
	app.Get("/contact", handlers.ContactHandler)
	///app.Get("/news", handlers.NewsHandler)

	app.Run(iris.Addr(":8090"), iris.WithCharset("UTF-8"))

}
