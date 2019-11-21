package providerlocations

import (
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryRestclientError(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	//Init
	// Execute
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get country BR", err.Message)
}

func TestGetCountryTimeout(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	//Init
	// Execute
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country  BR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("IR")
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/IR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
			"message": "Country not found",
			"error": "not_found",
			"status": 404,
			"cause": [
			]
		  }`,
	})
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)

}

func TestGetCountryInvalid(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country BR", err.Message)
}

func TestGetCountryInvalidJson(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal when getting country data for BR", err.Message)

}

func TestGetCountryNoError(t *testing.T) {
	rest.StartMockupServer()
	country, err := GetCountry("AR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.State))

}
