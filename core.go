package bri

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

const (
	TOKEN_PATH      = "/oauth/client_credential/accesstoken?grant_type=client_credentials"
	CREATE_VA_PATH  = "/v1/briva"
	BRI_TIME_FORMAT = "2006-01-02T15:04:05.999Z"
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

func (gateway *CoreGateway) CreateVA(token string, req CreateVaRequest) (res VaResponse, err error) {
	token = "Bearer " + token
	method := "POST"
	body, err := json.Marshal(req)
	timestamp := getTimestamp(BRI_TIME_FORMAT)
	signature := generateSignature(CREATE_VA_PATH, "POST", token, timestamp, string(body), gateway.Client.ClientSecret)

	headers := map[string]string{
		"Authorization": token,
		"BRI-Timestamp": timestamp,
		"BRI-Signature": signature,
		"Content-Type":  "application/json",
	}

	err = gateway.Call(method, CREATE_VA_PATH, headers, strings.NewReader(string(body)), &res)

	if err != nil {
		return
	}

	return
}
