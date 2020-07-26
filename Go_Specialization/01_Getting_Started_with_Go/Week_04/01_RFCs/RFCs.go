package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	// PROTOCOLS AND FORMATS

	// And RFC technically is a request for comments, but what it really
	// is, is a definition of a protocol or format. This is not specific
	// to Go, but when you make a big program, you'll want to interact
	// with other systems, other blocks of data somehow. This happens
	// with databases, network protocols, etc, there has to be a format
	// that everybody agrees on and everybody understand.

	// RFC-1866: HTML - Hyper Text Markup Language 2.0
	// RFC-3986: URI - Uniform Resource Identifier
	// RFC-2616: HTTP - Hyper Text Transfer Protocol
	// RFC-7159: JSON - JavaScript Object Notation

	// PROTOCOL PACKAGES

	// Go provides packages for important RFCs, functions which encode
	// and decode protocol format.

	// "net/http" : web communication protocol.
	stuff, err := http.Get("http://www.google.com")
	if err != nil {
		// If URL doesn't have "http://" it will throw this error:
		// Result: Get "www.google.com": unsupported protocol scheme ""
		fmt.Println(err)
	} else {
		// Result: *http.Response
		fmt.Printf("%T\n", stuff)
	}

	// JSON

	// Format to represent structured information, attribute-value pairs.
	// Can store in it a struct or a map.
	// Supports basic value types: Bool, number, string, array, "object".
	jsonResult, err := json.Marshal(struct {
		Name string `json:"name"`
		Mail string `json:"mail"`
		Age  int    `json:"age"`
	}{
		Name: "Santiago",
		Mail: "S@Z.com",
		Age:  28,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		// The marshalled JSON is a array of bytes...
		// Result: []uint8 [123 34 110 97 109 101 34 58 34 83 97 110 116 105 97 103 111 34 44 34 109 97 105 108 34 58 34 83 64 90 46 99 111 109 34 44 34 97 103 101 34 58 50 56 125]
		fmt.Printf("%T %v\n", jsonResult, jsonResult)
		// Result: {"name":"Santiago","mail":"S@Z.com","age":28}
		fmt.Printf("%v\n", string(jsonResult))
	}
}
