package povIDerlocations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"

	"gotrain/GoTestingInteg/mygolang-testing/api/domain/locations"
	"gotrain/GoTestingInteg/mygolang-testing/api/utils/errors"
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
			Message: fmt.Sprintf("InvalID response when trying to get country %s", countryID),
		}

	}
	var apiErr errors.APIerror
	if response.StatusCode > 299 {
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.APIerror{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("InvalID error response when getting country %s", countryID),
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
