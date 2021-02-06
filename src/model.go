package src

import "time"

func timeNow() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
}

// Request is a request incoming type
type Request struct {
	From        string `json:"from"`
	Message     string `json:"message"`
	Requesttime string `json:"requesttime"`
}

func newRequest(from string, message string) *Request {
	r := Request{from, message, timeNow()}
	return &r
}

func defaultRequest() *Request {
	return &Request{"", "", timeNow()}
}

// Response is a request outgoing type
type Response struct {
	To          string `json:"to"`
	Message     string `json:"message"`
	Respondtime string `json:"respondtime"`
}

func newResponse(to string, message string) *Response {
	r := Response{to, message, timeNow()}
	return &r
}
