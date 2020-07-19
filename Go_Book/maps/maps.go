package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// Vertex : a struct for coordinates in a map.
type Vertex struct {
	Lat, Long float64
}

// WordCount : count occurrences of words in a string.
func WordCount(s string) map[string]int {
	wordsMap := map[string]int{}
	words := strings.Split(s, " ")
	for _, word := range words {
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
		// can just do wordsMap[word]++ because then accessing
		// a nonexistent then the zero value will be 0, but this
		// way is better for cases where its something that needs
		// proper initialization.
	}
	return wordsMap
}

var m map[string]Vertex

func main() {
	// This only creates a nil map. Will behave as empty when reading, but
	// trying to write to it will cause a runtime error.
	// var m map[string]Vertex

	// Create a map either way.
	// m := map[string]Vertex{}
	m = make(map[string]Vertex)
	// Store a `Vertex` in the map `m` with the key "Bell Labs"
	k := "Bell Labs"
	m[k] = Vertex{
		Lat:  40.68433,
		Long: -74.39967,
	}
	// Can access the value from the key the same way.
	v := m[k]
	fmt.Printf("%T %v\n", v, v)

	// Maps can also be initialized with literals or empty.
	// The structs don't need literals, but the keys of the map does,
	// because then what would you be assigning to?
	m2 := map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Printf("%v\n", m2)

	// Delete the value of `k` from the map, but doesn't actually erase the
	// thing already taken out of the map and stored at `v`.
	delete(m2, k)
	fmt.Println(v, m2)
	// Now the value for `k` won't be in the map, so it will return a nil struct
	// { 0 0 } and a ok value to say if the struct was actually intended like that
	// or if its a nil value (not found).
	v, ok := m2[k]
	fmt.Printf("Not found and this is actually a nil struct,\t%v %v\n", v, ok)
	m2[k] = Vertex{0, 0}
	v, ok = m2[k]
	fmt.Printf("Found and this is an intended struct,\t\t%v %v\n", v, ok)

	// Doing a short initialization in an if checking if a value is in a map and
	// then only entering the if it indeed is.
	if value, ok := m2[k]; ok {
		fmt.Printf("Value %v is in the map!\n\n\n", value)
	}

	wc.Test(WordCount)
}
