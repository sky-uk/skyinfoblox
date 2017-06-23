package api

// RespError : what POST/PUT/DELETE requests returns in case of error.
type RespError struct {
	Error string `json:"Error"`
	Code  string `json:"code"`
	Text  string `json:"text"`
}
