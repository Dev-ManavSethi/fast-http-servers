//d fasthttp 1000 REQ 2277298
//D HTTP 1000 REQ 389030 e=42

//d fast http 10000 req 91416086 e=8800
//d http 10000 req 105629894 e=8800
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

func main() {

	color.Yellow("fast HTTP SERVER J")
	//http.HandleFunc("/", Home)

	// err := http.ListenAndServe(":8090", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	h := requestHandler
	if true {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(":9005", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Your ip is %s\n\n"+r.RemoteAddr)
	color.Green("Request Handled")

}

func requestHandler(ctx *fasthttp.RequestCtx) {

	fmt.Fprintf(ctx, "H")
	//fmt.Fprintf(ctx, "Hello, world!\n\n")

	// fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	// fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	// fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	// fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	// fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	// fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	// fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	// fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	// fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	//

	// fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	// ctx.SetContentType("text/plain; charset=utf8")

	// // Set arbitrary headers
	// ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// // Set cookies
	// var c fasthttp.Cookie
	// c.SetKey("cookie-name")
	// c.SetValue("cookie-value")
	// ctx.Response.Header.SetCookie(&c)
}
