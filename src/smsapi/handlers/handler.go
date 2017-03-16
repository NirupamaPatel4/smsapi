package handlers

import "github.com/valyala/fasthttp"

type Handler interface {
	Handle(ctx *fasthttp.RequestCtx)
}
