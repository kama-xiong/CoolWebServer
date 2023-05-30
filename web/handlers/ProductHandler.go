package handlers

import (
	"WebServer/services"
	"github.com/kataras/iris/v12"
)

func ProductPartyHandler(products iris.Party) {
	products.Get("/{bigcategory:string}/", mainCategoryHandler)
	products.Get("/{bigcategory:string}/{smallcategory:string}", smallCategoryHandler)

}

func mainCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	bigcategory := ctx.Params().Get("bigcategory")
	pagedata := ser.GetCategoriesByName(bigcategory)
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
func smallCategoryHandler(ctx iris.Context) {
	ser := services.CategoryService{}
	bigcategory := ctx.Params().Get("bigcategory")
	smallcategory := ctx.Params().Get("smallcategory")
	pagedata := ser.GetCategoriesByName(bigcategory, smallcategory)
	ctx.ViewData("Model", *pagedata)
	ctx.View("/products/paperbox.html")
}
