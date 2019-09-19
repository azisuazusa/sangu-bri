package bri

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiredTime string   `json:"expires_in"`
	ProductList []string `json:"api_product_list_json"`
}
