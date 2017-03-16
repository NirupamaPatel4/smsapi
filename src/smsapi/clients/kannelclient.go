package clients

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	KANNEL_USER string = "bar"
	KANNEL_PWD  string = "bar"
	KANNEL_SMSC string = "SMPPSim"
	KANNEL_HOST string = "localhost"
	KANNEL_PORT string = "14010"
)

type KannelClient struct {
	baseUri string
}

func (kannelClient *KannelClient) SendSms(from string, to string, text string) (status int, err error) {

	uri := fmt.Sprintf("%s&from=%s&to=%s&text=%s", kannelClient.baseUri, from, to, text)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err = client.Do(req, resp)

	bodyBytes := resp.Body()
	status = resp.StatusCode()
	fmt.Println(string(bodyBytes))

	return
}

func NewKannelClient() *KannelClient {
	uri := fmt.Sprintf("http://%s:%s/cgi-bin/sendsms?username=%s&password=%s&smsc=%s",
		KANNEL_HOST, KANNEL_PORT, KANNEL_USER, KANNEL_PWD, KANNEL_SMSC)
	return &KannelClient{uri}
}
