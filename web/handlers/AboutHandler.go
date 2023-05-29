package handlers

import "github.com/kataras/iris/v12"

func AboutHandler(ctx iris.Context) {
	ctx.View("homePage.html")
}
