package bri

import (
	"strings"
	"testing"

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
