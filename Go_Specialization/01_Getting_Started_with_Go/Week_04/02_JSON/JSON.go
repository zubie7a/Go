package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// JSON

	// Format to represent structured information, attribute-value pairs.
	// Can store in it a struct or a map. It's all represented as unicode
	// characters, generally human readable, and its representation is
	// fairly compact.

	// Supports basic value types: Bool, number, string, array, "object".
	// Types can be combined recursively: arrays of structs, struct in
	// struct, arrays of certain basic  types, etc.

	jsonResult, err := json.Marshal(struct {
		// An "unexported" member of the struct won't be detected in the
		// JSON function and won't be encoded.
		city string
		Name string // Without JSON anotation will use the field name.
		Mail string `json:"email"` // JSON annotations can be given for
		Age  int    `json:"age"`   // names that will be used in the JSON.
	}{
		city: "Amsterdam", // This won't be in JSON.
		Name: "Santiago",  // This will be in JSON with field name "Name".
		Mail: "S@Z.com",   // This will be in JSON with field name "email".
		Age:  28,          // This will be in JSON with field name "age".
	})

	if err != nil {
		fmt.Println(err)
	} else {
		// The marshalled JSON is a array of bytes...
		// []uint8 [123 34 78 97 109 101 34 58 34 83 97 110 116 105 97 103 111 34 44 34 101 109 97 105 108 34 58 34 83 64 90 46 99 111 109 34 44 34 97 103 101 34 58 50 56 125]
		fmt.Printf("%T %v\n", jsonResult, jsonResult)
		// {"Name":"Santiago","email":"S@Z.com","age":28}
		fmt.Printf("%v\n", string(jsonResult))
	}

	var p2 struct {
		city string
		Name string
		Mail string `json:"email"`
		// Age  int    `json:"age"`
		Cash int `json:"cash"`
	}

	err = json.Unmarshal(jsonResult, &p2)
	if err == nil {
		// Fields not in the JSON won't be unmarshalled.
		p2.city = "Amsterdam"
		// There's no "age" because the struct doesn't have it, even if the
		// JSON does have it, so it is ignored.
		// Cash will have default value 0 because nothing in the JSON matches.
		// Result: {Amsterdam Santiago S@Z.com 0}
		fmt.Println(p2)
	}

}
