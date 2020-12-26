package entities

// Response define data structure for http response object
type Response struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Value     int    `json:"value"`
}
