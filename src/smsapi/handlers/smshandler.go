package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"smsapi/clients"
	"smsapi/models"
)

type SMSHandler struct {
	smsClient clients.SMSClientInterface
}

func (smsHandler *SMSHandler) Handle(ctx *fasthttp.RequestCtx) {
	isValid, smsRequest, err := smsHandler.validateAndCreateSMSRequest(ctx)
	if isValid == false {
		if err == nil {
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusBadRequest), fasthttp.StatusBadRequest)
		} else {
			response := &models.SMSResponse{
				Error: err.Error(),
			}
			r, _ := json.Marshal(response)
			ctx.Error(string(r), fasthttp.StatusBadRequest)
		}
		return
	} else {
		status, senderr := smsHandler.smsClient.SendSms(smsRequest.From, smsRequest.To, smsRequest.Text)
		var response *models.SMSResponse
		if status == 202 && senderr == nil {
			response = &models.SMSResponse{
				Message: "outbound sms ok",
			}
			r, _ := json.Marshal(response)
			ctx.Success("application/json", r)
		} else {
			response = &models.SMSResponse{
				Error: "Unexpected error occurred",
			}
			r, _ := json.Marshal(response)
			ctx.Error(string(r), fasthttp.StatusInternalServerError)
		}
	}
}

func GetSMSHandlerForTest(smsClient clients.SMSClientInterface) *SMSHandler {
	return &SMSHandler{smsClient}
}

func GetSMSHandler() *SMSHandler {
	return &SMSHandler{clients.NewKannelClient()}
}

func (smsHandler *SMSHandler) validateAndCreateSMSRequest(ctx *fasthttp.RequestCtx) (isValid bool, smsRequest models.SMSRequest, err error) {
	isValid = false
	smsRequest = models.SMSRequest{}
	err = nil

	if contentType := string(ctx.Request.Header.ContentType()); contentType != "application/json" {
		return
	}

	postBody := ctx.PostBody()
	if len(postBody) == 0 {
		return
	}

	json.Unmarshal(postBody, &smsRequest)
	err = smsRequest.Validate()
	if err == nil {
		isValid = true
	}
	return
}
