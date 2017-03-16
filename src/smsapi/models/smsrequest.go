package models

import (
	"errors"
	"gopkg.in/validator.v2"
)

type SMSRequest struct {
	From string `json:"from" validate:"nonzero,min=6,max=16,regexp=^[0-9]*$"`
	To   string `json:"to" validate:"nonzero,min=6,max=16,regexp=^[0-9]*$"`
	Text string `json:"text" validate:"nonzero,min=1,max=120"`
}

func (smsRequest *SMSRequest) Validate() error {
	err := validator.Validate(smsRequest)
	if err == nil {
		return err
	}
	errs, _ := err.(validator.ErrorMap)
	errorString := ""
	for k, v := range errs {
		errorString += k
		if v[0] == validator.ErrZeroValue {
			errorString += " is missing."
		} else {
			errorString += " is invalid."
		}
	}
	return errors.New(errorString)
}
