package handlers

import (
	"errors"
	"github.com/valyala/fasthttp"
	"testing"
)

type SuccesfulSMSClient struct {
}

func (smsClient *SuccesfulSMSClient) SendSms(from string, to string, text string) (status int, err error) {
	status = 202
	err = nil
	return
}

type FailedSMSClient struct {
}

func (smsClient *FailedSMSClient) SendSms(from string, to string, text string) (status int, err error) {
	status = 500
	err = errors.New("Something bad happened")
	return
}

func TestHandleInvalidContentType(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.Header.SetContentType("application/text")
	req.SetBodyString("{ \"From\" : \"123456\", \"To\": \"123456\", \"Text\": \"Hi\"}")
	ctx.Init(&req, nil, nil)
	GetSMSHandlerForTest(&SuccesfulSMSClient{}).Handle(&ctx)

	if ctx.Response.StatusCode() != 400 {
		t.Fatalf("unexpected status code %d. Expecting 400", ctx.Response.StatusCode())
	}

	body := string(ctx.Response.Body())
	expected := "Bad Request"
	if body != expected {
		t.Fatalf("unexpected status message %s. Expecting %s", body, expected)
	}
}

func TestHandleEmptyBody(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.Header.SetContentType("application/json")
	req.SetBodyString("")
	ctx.Init(&req, nil, nil)
	GetSMSHandlerForTest(&SuccesfulSMSClient{}).Handle(&ctx)

	if ctx.Response.StatusCode() != 400 {
		t.Fatalf("unexpected status code %d. Expecting 400", ctx.Response.StatusCode())
	}

	body := string(ctx.Response.Body())
	expected := "Bad Request"
	if body != expected {
		t.Fatalf("unexpected status message %s. Expecting %s", body, expected)
	}
}

func TestHandleInvalidBody(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.Header.SetContentType("application/json")
	req.SetBodyString("{ \"Wrong\" : \"123456\", \"To\": \"123456\", \"Text\": \"Hi\"}")
	ctx.Init(&req, nil, nil)
	GetSMSHandlerForTest(&SuccesfulSMSClient{}).Handle(&ctx)

	if ctx.Response.StatusCode() != 400 {
		t.Fatalf("unexpected status code %d. Expecting 400", ctx.Response.StatusCode())
	}

	body := string(ctx.Response.Body())
	expected := "{\"message\":\"\",\"error\":\"From is missing.\"}"
	if body != expected {
		t.Fatalf("unexpected status message %s. Expecting %s", body, expected)
	}
}

func TestHandleValidRequest(t *testing.T) {
	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.Header.SetContentType("application/json")
	req.SetBodyString("{ \"From\" : \"123456\", \"To\": \"123456\", \"Text\": \"Hi\"}")
	ctx.Init(&req, nil, nil)
	GetSMSHandlerForTest(&SuccesfulSMSClient{}).Handle(&ctx)

	if ctx.Response.StatusCode() != 200 {
		t.Fatalf("unexpected status code %d. Expecting 200", ctx.Response.StatusCode())
	}

	body := string(ctx.Response.Body())
	expected := "{\"message\":\"outbound sms ok\",\"error\":\"\"}"
	if body != expected {
		t.Fatalf("unexpected status message %s. Expecting %s", body, expected)
	}
}

func TestHandleSMSClientError(t *testing.T) {

	var ctx fasthttp.RequestCtx
	var req fasthttp.Request
	req.Header.SetContentType("application/json")
	req.SetBodyString("{ \"From\" : \"123456\", \"To\": \"123456\", \"Text\": \"Hi\"}")
	ctx.Init(&req, nil, nil)
	GetSMSHandlerForTest(&FailedSMSClient{}).Handle(&ctx)

	if ctx.Response.StatusCode() != 500 {
		t.Fatalf("unexpected status code %d. Expecting 500", ctx.Response.StatusCode())
	}

	body := string(ctx.Response.Body())
	expected := "{\"message\":\"\",\"error\":\"Unexpected error occurred\"}"
	if body != expected {
		t.Fatalf("unexpected status message %s. Expecting %s", body, expected)
	}
}
