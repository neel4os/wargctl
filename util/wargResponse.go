package util

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type methods interface {
	Post(path string, payload map[string]interface{})
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

type Org struct {
	name               string
	description        string
	organizationId     string
	creation_time      string
	last_modified_time string
	version            int
}

func (w *wargCtlResponse) Post(path string, body map[string]interface{}) {
	resp, error := w.client.R().SetBody(body).Post(path)
	fmt.Println(resp, resp.IsSuccess(), resp.StatusCode(), resp.String())
	if resp.IsSuccess() {
		TablePrinter{resp.Body()}.Print()
	} else {
		fmt.Println(error)
	}

}
