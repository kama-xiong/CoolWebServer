package handlers

import (
	"WebServer/services"
	"github.com/kataras/iris/v12"
)

func ProductPartyHandler(products iris.Party) {
	products.Get("/paperbox/", mainPaperboxCategoryHandler)
	products.Get("/paperbox/{name:string}", smallPaperboxCategoryHandler)

}

func mainPaperboxCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	pagedata := ser.GetCategoriesByName("paper box")
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
func smallPaperboxCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	name := ctx.Params().Get("name")
	pagedata := ser.GetCategoriesByName(name)
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
func mainCoveredboxCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	pagedata := ser.GetCategoriesByName("covered box")
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
func smallCoveredCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	name := ctx.Params().Get("name")
	pagedata := ser.GetCategoriesByName(name)
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
