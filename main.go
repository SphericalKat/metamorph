package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
)

const indexString = "metamorph v0.1\ngithub.com/ATechnoHazard/metamorph"
var ok = []byte{'O', 'K'}

func index(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(indexString)
}

func alive(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.Write(ok)
}

func main() {
	mux := router.New()
	mux.PanicHandler = func(ctx *fasthttp.RequestCtx, i interface{}) {
		ctx.SetStatusCode(http.StatusInternalServerError)
		log.Println("PANIC:", i)
	}

	base := os.Getenv("BASE_PATH")

	mux.GET(base+"/", index)
	mux.GET(base+"/alive", alive)

	log.Println("Starting metamorph")
	log.Fatal(fasthttp.ListenAndServe(":8081", mux.Handler))
}
