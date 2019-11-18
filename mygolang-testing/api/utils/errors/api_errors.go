package errors

type APIerror struct {
	Status  int    `json: "status"`
	Message string `json: "message"`
	Error   string `json: "error"`
}
