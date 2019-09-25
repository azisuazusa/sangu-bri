package bri

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BriSanguTestSuite struct {
	suite.Suite
}

func TestBriSanguTestSuite(t *testing.T) {
	suite.Run(t, new(BriSanguTestSuite))
}

func (bri *BriSanguTestSuite) TestGetTokenSuccess() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"

	coreGateway := CoreGateway{
		Client: briClient,
	}

	resp, err := coreGateway.GetToken()
	containsProduct := false
	for _, v := range resp.ProductList {
		if strings.Contains(v, "briva") {
			containsProduct = true
			break
		}
	}

	assert.NotNil(bri.T(), resp.AccessToken)
	assert.Equal(bri.T(), "179999", resp.ExpiredTime)
	assert.Equal(bri.T(), true, containsProduct)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidKeySecret() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "abcdef"
	briClient.ClientSecret = "123"

	coreGateway := CoreGateway{
		Client: briClient,
	}

	resp, err := coreGateway.GetToken()

	assert.Equal(bri.T(), "", resp.AccessToken)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidUrl() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id/v1"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"

	coreGateway := CoreGateway{
		Client: briClient,
	}

	resp, err := coreGateway.GetToken()

	assert.NotNil(bri.T(), err)
	assert.Equal(bri.T(), "", resp.AccessToken)
}

func (bri *BriSanguTestSuite) TestGetTokenFailedInvalidProduct() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "QYTaZS7Rw3JztRWhXAYrXjKUAg13AvRa"
	briClient.ClientSecret = "UiqE9zGXcPd6iXGv"

	coreGateway := CoreGateway{
		Client: briClient,
	}

	resp, err := coreGateway.GetToken()

	containsProduct := false
	for _, v := range resp.ProductList {
		if strings.Contains(v, "briva") {
			containsProduct = true
			break
		}
	}

	assert.Equal(bri.T(), false, containsProduct)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaSuccess() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaFailedDuplicate() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreateVA(token, req)

	// create second request
	resp, err = coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "13", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestCreateVaFailedExpiredMoreThanThreeMonths() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "123123" + random,
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.CreateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "12", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaSuccess() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 0)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaFailedCustomerNotFound() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313555",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "14", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestUpdateVaFailedExpiredMoreThanThreeMonths() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	random := strconv.Itoa(rand.Intn(10000))
	dt := time.Now().AddDate(0, 3, 1)
	expired := dt.Format("2006-01-02 15:04:05")
	req := CreateVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		CustCode:        "1231233313",
		Name:            "Orang Baik " + random,
		Amount:          random,
		Description:     "test",
		ExpiredDate:     expired,
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.UpdateVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "12", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaSuccess() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190918",
		EndDate:         "20190918",
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), true, resp.Status)
	assert.Equal(bri.T(), "00", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaFailedNoTransaction() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190919",
		EndDate:         "20190919",
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "41", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}

func (bri *BriSanguTestSuite) TestGetReportVaFailedInvalidDateRange() {
	briClient := NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	coreGateway := CoreGateway{
		Client: briClient,
	}
	tokenResp, err := coreGateway.GetToken()

	req := GetReportVaRequest{
		InstitutionCode: "J104408",
		BrivaNo:         "77777",
		StartDate:       "20190917",
		EndDate:         "20190918",
	}

	token := tokenResp.AccessToken
	resp, err := coreGateway.GetReportVA(token, req)

	assert.Equal(bri.T(), false, resp.Status)
	assert.Equal(bri.T(), "42", resp.ResponseCode)
	assert.Equal(bri.T(), nil, err)
}
