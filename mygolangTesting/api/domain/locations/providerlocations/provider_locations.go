package providerlocations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"

	"gotrain/GoTestingInteg/mygolangTesting/api/domain/locations"
	"gotrain/GoTestingInteg/mygolangTesting/api/utils/errors"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

// GetCountry takes a string and returns Country data
func GetCountry(countryID string) (*locations.Country, *errors.APIerror) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryID))
	if response == nil || response.Response == nil {
		return nil, &errors.APIerror{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryID),
		}

	}

	if response.StatusCode > 299 {
		var apiErr errors.APIerror
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.APIerror{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryID),
			}
		}
		return nil, &apiErr
	}
	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.APIerror{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal when getting country data for %s", countryID),
		}

	}

	return &result, nil
}
