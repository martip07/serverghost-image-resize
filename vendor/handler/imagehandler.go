package handler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Howdy API Image Resize")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Path:, %s\n", ctx.UserValue("filepath"))
	// fmt.Printf(ctx.QueryArgs().String())
	// var queryValue = ctx.QueryArgs().Peek("width")
	if ctx.QueryArgs().Has("width") {
		fmt.Println(string(ctx.QueryArgs().Peek("width")))
	}
	// fmt.Println(string(queryValue))
}
