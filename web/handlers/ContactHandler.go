package handlers

import "github.com/kataras/iris/v12"

func ContactHandler(ctx iris.Context) {
	ctx.View("/contact/contact.html")
}
