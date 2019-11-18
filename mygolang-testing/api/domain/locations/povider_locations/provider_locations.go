package povider_locations

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"gotrain/GoTestingInteg/mygolang-testing/api/domain/locations"
	"gotrain/GoTestingInteg/mygolang-testing/api/utils/errors"
	"net/http"
	"net/url"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.CountryId, *errors.APIerror) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	if response == nil || response.Response == nil {
		return nil, errors.APIerror{
			Status:  http.StatusInternalServerError,
			Message: Sprintf("Invalid response when trying to get country %s", countryId),
		}

	}
	var apiErr errors.APIerror
	if response.StatusCode > 299 {
     if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
		 return nil, &errors.APIerror{
			 Status: http.StatusInternalServerError,
			 Message: fmt.Sprintf("Invalid error response when getting country %", countryId)
		 }
		 return nil, &apiErr
	 }
	 var result locations.Country
	 if

	}
	return nil, nil
}
