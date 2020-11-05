package main

import (
	"encoding/json"
	"fmt"
	//	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"pages"`
	Fruits []string `json:"fruits"`
}

// Marshal custom data type
func MarshalCustomType() {
	fmt.Println("Marshal a custom type without using field tag...")
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	// res1B is a slice of bytes []byte
	res1B, _ := json.Marshal(res1D)
	fmt.Println("\n", string(res1B))
	fmt.Printf("\nType of res1B: %T\n", res1B) // []uint8 (alias of byte)

	fmt.Println("\nMarshal a custom type with using filed tag...")
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	// The key of json object becomes small: e.g. Page -> pages (defined in field tag)
	fmt.Println("\n", string(res2B))
}

func Unmarshal() {
	// Data to be unmarshaled
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// A variable to hold the decoded data
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("Unmarshaling...")
	fmt.Println("Input: ", byt)
	fmt.Println("Output: ", dat)
	fmt.Printf("data type of dat? %T\n", dat)

	fmt.Println("\nLoop throught all key-value pairs in unmarshaled data...")
	for k, _ := range dat {
		if v, ok := dat[k]; ok {
			fmt.Printf("keys: %s, values: %v\n", k, v)
		}
	}

	// What does (float64) do?
	fmt.Println("\nCheck value of json object key... ")
	num := dat["num"].(float64)
	fmt.Println("value of 'num':", num)

	// Access nested data needs conversion
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("\nvalue of key 'strs': ", str1)

}

// Marshal basic data type
func MarshalBasicType() {
	fmt.Println("Marshal a bool...")
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	fmt.Println("\nMarshal a int...")
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fmt.Println("\nMarshal a float...")
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	fmt.Println("\nMarshal a string...")
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	fmt.Println("\nMarshal a slice of string...")
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	fmt.Println("\nMarshal a map of string to int...")
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

}
func main() {
	MarshalBasicType()
	fmt.Println("----------------------------")
	MarshalCustomType()
	fmt.Println("----------------------------")
	Unmarshal()
}
