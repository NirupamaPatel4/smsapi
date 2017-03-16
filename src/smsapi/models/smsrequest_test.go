package models

import (
	"testing"
)

func TestValidateRequiredField(t *testing.T) {
	smsRequest := &SMSRequest{
		From: "123456",
		Text: "Hi",
	}
	expected := "To is missing."
	actual := smsRequest.Validate().Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestValidateMultipleRequiredFields(t *testing.T) {
	smsRequest := &SMSRequest{
		Text: "Hi",
	}
	expected := "From is missing.To is missing."
	actual := smsRequest.Validate().Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestValidateInvalidFormatInFrom(t *testing.T) {
	smsRequest := &SMSRequest{
		From: "shilpa",
		To:   "123456",
		Text: "Hi",
	}
	expected := "From is invalid."
	actual := smsRequest.Validate().Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestValidateInvalidLengthInTo(t *testing.T) {
	smsRequest := &SMSRequest{
		From: "987867",
		To:   "12345",
		Text: "Hi",
	}
	expected := "To is invalid."
	actual := smsRequest.Validate().Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestValidateInvalidText(t *testing.T) {
	smsRequest := &SMSRequest{
		From: "987867",
		To:   "123456",
		Text: "The Ides of March (Latin: Idus Martiae, Late Latin: Idus Martii) is a day on the Roman calendar that corresponds to March 15. It was marked by several religious observances and became notorious as the date of the assassination of Julius Caesar in 44 BC.",
	}
	expected := "Text is invalid."
	actual := smsRequest.Validate().Error()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestValidateValidRequest(t *testing.T) {
	smsRequest := &SMSRequest{
		From: "123456",
		To:   "123456",
		Text: "Hi",
	}
	actual := smsRequest.Validate()
	if actual != nil {
		t.Errorf("Test failed, expected: nil, got:  '%s'", actual)
	}
}
