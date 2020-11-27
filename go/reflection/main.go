package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println("Type: ", t)
	fmt.Println("Kind: ", v.Kind())
	fmt.Println("Value: ", v)

	// Show the field of q if q is struct
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields: ", v.NumField())

		// Show the fields of struct
		for i := 0; i < v.NumField(); i++ {
			fmt.Println("Field: ", v.Field(i), "Type:", v.Field(i).Type())
		}
	}
}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}
