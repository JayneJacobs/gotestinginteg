package povider_locations

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetCountryTimeout(t *testing.T) {

	//Init
	// Execute
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid restclient response when getting country, BR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	country, err := GetCountry("IR")
	//Validate
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)

}

func TestGetCountryInvalid(t *testing.T) {
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal when getting country data for, BR", err.Message)
}

func TestGetCountryInvalidJson(t *testing.T) {
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal when getting country data for, BR", err.Message)

}

func TestGetCountryNoError(t *testing.T) {
	country, err := GetCountry("AR")
	//Validate
	assert.Nil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT -03:00", country.TimeZone)
	assert.EqualValues(t, "error when trying to unmarshal when getting country data for, BR", err.Message)

}
