package router

import (
	"github.com/valyala/fasthttp"
	"os"
	"testing"
)

func TestHandleGetToOutboundSmsRoute(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://localhost:8080/outbound/sms/")
	ctx.Init(&req, nil, nil)
	HandleRoute(&ctx)

	if ctx.Response.StatusCode() != 405 {
		t.Fatalf("unexpected status code %d. Expecting 405", ctx.Response.StatusCode())
	}
	body := string(ctx.Response.Body())
	if body != "Method Not Allowed" {
		t.Fatalf("unexpected status message %s. Expecting Method Not Allowed", body)
	}
}

func TestHandleInvalidRoute(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://localhost:8080/outbound/call/")
	ctx.Init(&req, nil, nil)
	HandleRoute(&ctx)

	if ctx.Response.StatusCode() != 404 {
		t.Fatalf("unexpected status code %d. Expecting 404", ctx.Response.StatusCode())
	}
	body := string(ctx.Response.Body())
	if body != "Not Found" {
		t.Fatalf("unexpected status message %s. Expecting Not Found", body)
	}
}

func TestHandleUnAuthenticatedPostToOutboundSmsRoute(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://localhost:8080/outbound/sms/")
	req.Header.SetMethod("POST")
	ctx.Init(&req, nil, nil)
	HandleRoute(&ctx)

	if ctx.Response.StatusCode() != 401 {
		t.Fatalf("unexpected status code %d. Expecting 401", ctx.Response.StatusCode())
	}
	body := string(ctx.Response.Body())
	if body != "Unauthorized" {
		t.Fatalf("unexpected status message %s. Expecting Unauthorized", body)
	}
}

func TestHandlePostToOutboundSmsRouteWithInvalidCredentials(t *testing.T) {
	os.Setenv("SMSAPI_PASS", "")
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://localhost:8080/outbound/sms/")
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Basic c2hpbHBhOa==")
	ctx.Init(&req, nil, nil)
	HandleRoute(&ctx)

	if ctx.Response.StatusCode() != 403 {
		t.Fatalf("unexpected status code %d. Expecting 403", ctx.Response.StatusCode())
	}
	body := string(ctx.Response.Body())
	if body != "Forbidden" {
		t.Fatalf("unexpected status message %s. Expecting Forbidden", body)
	}
}

func TestHandlePostToOutboundSmsRouteWithValidCredentials(t *testing.T) {
	os.Setenv("SMSAPI_PASS", "")
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.SetRequestURI("http://localhost:8080/outbound/sms/")
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", "Basic c2hpbHBhOg==")
	req.Header.SetContentType("application/json")
	req.SetBodyString("{ \"From\" : \"123456\", \"To\": \"123456\", \"Text\": \"Hi\"}")
	ctx.Init(&req, nil, nil)
	HandleRoute(&ctx)

	if ctx.Response.StatusCode() != 200 {
		t.Fatalf("unexpected status code %d. Expecting 200", ctx.Response.StatusCode())
	}
}
