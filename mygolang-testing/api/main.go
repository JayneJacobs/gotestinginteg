package main

import (
	"fmt"

	"src/gotrain/GoTestingInteg/mygolang-testing/api/domain/locations/providerlocations"
)

func main() {
	country, err := providerlocations.GetCountries("ARS")
	fmt.Println(err)
	fmt.Println(country)
}
