package handler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Howdy API Image Resize")
}

func Resize(ctx *fasthttp.RequestCtx) {

	// fmt.Printf(ctx.QueryArgs().String())
	// var queryValue = ctx.QueryArgs().Peek("width")
	var pathImage string
	var widthImage string
	var heightImage string
	if ctx.UserValue("filepath") != "/" {
		pathImage = fmt.Sprint(ctx.UserValue("filepath"))
		fmt.Println(pathImage)
	}
	if ctx.QueryArgs().Has("width") {
		widthImage = string(ctx.QueryArgs().Peek("width"))
		fmt.Println(widthImage)
	}
	if ctx.QueryArgs().Has("height") {
		heightImage = string(ctx.QueryArgs().Peek("height"))
		fmt.Println(heightImage)
	}
	ctx.WriteString(pathImage + widthImage + heightImage)
	//fmt.Fprintf(ctx, "Path:, %s\n", pathImage)
	// fmt.Println(string(queryValue))
}
