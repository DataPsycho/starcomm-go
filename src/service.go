package src

import "fmt"

func reply(r *Request) *Response {
	message := fmt.Sprintf("Listening to %s, from Discovery!", r.From)
	response := newResponse(r.From, message)
	return response
}
