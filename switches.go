package main

type HTTPRequest struct {
	Method string
}

func switches() {

	r := HTTPRequest{Method: "HEAD"}

	// Implicit breaking, no "break" command needed, but you need to fall through use "fallthrough"
	switch r.Method {
	case "GET":
		println("GET request")
		fallthrough
	case "DELETE":
		println("DELETE request")
	case "POST":
		println("POST request")
	case "PUT":
		println("PUT request")
	default:
		println("Unhandled method")
	}

}
