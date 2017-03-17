package router

import (
	"bytes"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	"smsapi/handlers"
)

func HandleRoute(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/outbound/sms/":
		if !ctx.IsPost() {
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed), fasthttp.StatusMethodNotAllowed)
			return
		}
		basicAuth(ctx, handlers.GetSMSHandler())
	default:
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
	}
}

var basicAuthPrefix = []byte("Basic ")

// BasicAuth is the basic auth handler
func basicAuth(ctx *fasthttp.RequestCtx, next handlers.Handler) {

	user := []byte("shilpa")
	pass := []byte("821bc092-fcc1-4e3c-8420-3357fd1b36e6")
	// Get the Basic Authentication credentials
	auth := ctx.Request.Header.Peek("Authorization")
	if bytes.HasPrefix(auth, basicAuthPrefix) {
		// Check credentials
		payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
		if err == nil {
			pair := bytes.SplitN(payload, []byte(":"), 2)
			if len(pair) == 2 &&
				bytes.Equal(pair[0], user) &&
				bytes.Equal(pair[1], pass) {
				// Delegate request to the given handle
				next.Handle(ctx)
				return
			} else {
				ctx.Error(fasthttp.StatusMessage(fasthttp.StatusForbidden), fasthttp.StatusForbidden)
				return
			}
		}
	}
	// Request Basic Authentication otherwise
	ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
}
