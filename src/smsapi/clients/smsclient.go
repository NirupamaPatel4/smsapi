package clients

type SMSClientInterface interface {
	SendSms(from string, to string, text string) (status int, err error)
}
