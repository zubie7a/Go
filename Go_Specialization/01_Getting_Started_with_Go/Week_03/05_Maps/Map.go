package main

import "fmt"

func main() {

	// MAPS

	// In Go, Maps are an implementation of a hash table.
	// Create a map that maps from string to int.
	// The syntax is `map[key_type]value_type`.
	var idMap map[string]int
	idMap = make(map[string]int)

	idMap["StringKey1"] = 7777777
	idMap["StringKey2"] = 2222222
	idMap["StringKey3"] = 1111111
	// Result: map[StringKey1:7777777 StringKey2:2222222 StringKey3:1111111]
	fmt.Printf("%v\n", idMap)

	// You can also initialize a map with some values.
	stuffMap := map[string]int{
		"SK1": 7,
		"SK2": 2,
		"SK3": 1,
		"SK4": 10, // Yay you can have trailing commas :-)
	}
	// Add another value after the initialization.
	stuffMap["SK0"] = 0
	// Delete one of the previous values.
	delete(stuffMap, "SK4")
	// Deleting a non-existent key won't cause trouble.
	delete(stuffMap, "SK5")
	// Result: map[SK0:0 SK1:7 SK2:2 SK3:1]
	fmt.Println(stuffMap)

	// Result: map[string]int
	mapType := fmt.Sprintf("%T", stuffMap)
	fmt.Println(mapType)

	// Uhh... when querying a map for a key that doesn't exist, we don't get
	// null... we get the default value for the type, which in the case of
	// `int` is 0.
	nonExistent := stuffMap["SKZ"]
	// Result: 0
	fmt.Println(nonExistent)
	// But 0 can be a valid value! For example, the value of "SK0". So how
	// do we distinguish between zero-value entries and non-existent entries?
	// We can do two-value assigment when querying a map. The second value `p`
	// is going to be a boolean, telling if the key is present in the map.
	testKey := "SKZ"
	testVal, p := stuffMap[testKey]
	if p {
		fmt.Printf("Key %s exists with value %d!\n", testKey, testVal)
	} else {
		fmt.Printf("Key %s doesn't exists!\n", testKey)
	}

	// If you want to know the amount of values in map...
	// Result: 4
	fmt.Println(len(stuffMap))

	for k, v := range stuffMap {
		// Iterating through a map won't give the same order than the keys
		// were inserted, this map is "unordered".
		// Result:
		// 	SK0 : 0
		// 	SK3 : 1
		// 	SK2 : 2
		// 	SK1 : 7
		fmt.Printf("%s : %d\n", k, v)
	}

	// There's no "clear" method to get the keys from a map, you have to
	// iterate yourself and build the list of keys.
	keys := make([]string, 0, len(stuffMap))
	for k := range stuffMap {
		keys = append(keys, k)
	}
	// Result: [SK3 SK0 SK1 SK2]
	fmt.Println(keys)
}
