package bri

import (
	"io"
	"net/url"
	"strings"
)

const (
	TOKEN_PATH = "oauth/client_credential/accesstoken?grant_type=client_credentials"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) GetToken() (res TokenResponse, err error) {
	data := url.Values{}
	data.Set("client_id", gateway.Client.ClientId)
	data.Set("client_secret", gateway.Client.ClientSecret)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", TOKEN_PATH, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}
