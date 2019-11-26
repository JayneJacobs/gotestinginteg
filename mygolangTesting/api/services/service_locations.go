package services

import (
	"fmt"
	"gotrain/GoTestingInteg/mygolangTesting/api/domain/locations"
	"gotrain/GoTestingInteg/mygolangTesting/api/domain/locations/providerlocations"
	"gotrain/GoTestingInteg/mygolangTesting/api/utils/errors"
)

// initialized = false set to true when the package is called
type locationsService struct{}

type locationServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.APIerror)
}

// LocationsService is the locationServiceInterface
var LocationsService locationServiceInterface

func init() {
	fmt.Println("Init Service")
	LocationsService = &locationsService{}
}

// GetCountry takes a countryId
// and returns a pointer to locations.Country and a pointer to an APIerror
func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.APIerror) {
	fmt.Println("Inside Service")
	return providerlocations.GetCountry(countryId)

}