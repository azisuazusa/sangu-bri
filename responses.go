package bri

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiredTime string   `json:"expires_in"`
	ProductList []string `json:"api_product_list_json"`
}

type VaResponse struct {
	Status              bool   `json:"status"`
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
	ErrDesc             string `json:"errDesc"`
	Data                VaData `json:"data"`
}

type VaData struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Nama            string `json:"nama"`
	Amount          string `json:"amount"`
	Keterangan      string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}
