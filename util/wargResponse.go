package util

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type methods interface {
	Post(path string, payload map[string]interface{})
	Get(path string)
}

var instantiated *wargCtlResponse = nil

type wargCtlResponse struct {
	//HostUrl string
	client *resty.Client
}

func NewWargCtlResponse() *wargCtlResponse {
	if instantiated == nil {
		wargClClient := resty.New()
		wargClClient.SetHeader("Content-Type", "application/json")
		instantiated = &wargCtlResponse{wargClClient}
	}
	return instantiated
}
func (w *wargCtlResponse) SetBaseUrl(baseurl string) *wargCtlResponse {
	w.client.SetBaseURL(baseurl)
	return w
}

func (w *wargCtlResponse) Post(path string, body map[string]interface{}) {
	resp, error := w.client.R().SetBody(body).Post(path)
	if resp.StatusCode() > 0 {
		TablePrinter{resp.Body()}.Print()
	} else {
		fmt.Println(error)
	}

}

func (w wargCtlResponse) Get(path string) {
	resp, err := w.client.R().Get(path)
	if resp.StatusCode() > 0 {
		TablePrinter{resp.Body()}.Print()
	} else {
		fmt.Println(err)
	}

}
