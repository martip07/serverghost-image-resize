package http

import (
	"log"

	handlerHTTP "handler"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func ImageHttp() {
	r := router.New()
	r.GET("/", handlerHTTP.Index)
	r.GET("/hello/*filepath", handlerHTTP.Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
