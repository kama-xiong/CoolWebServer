package handlers

import (
	"github.com/kataras/iris/v12"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func NewsHandler(ctx iris.Context) {
	input := []byte(`# 5lmh.com是个不错的go文档网站
\![Alt](/static/images/category/pouch.jpg)
\<img src="/static/images/category/display.jpg" width="200px"/>`)
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	ctx.ViewData("Content", string(html))
	ctx.View("/news/NewsPage.html")

}
