package app

import "gotrain/GoTestingInteg/mygolangTesting/api/controllers"

func mapURLs() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
