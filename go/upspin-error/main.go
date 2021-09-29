package main

import (
	"fmt"

	"upspin.io/errors"
	"upspin.io/upspin"
)

func main() {
	path := upspin.PathName("jane@doe.com/file")
	user := upspin.UserName("joe@blow.com")

	// Single error.
	e1 := errors.E(errors.Op("Get"), path, errors.IO, "network unreachable")
	fmt.Println("\nSimple error:")
	fmt.Println(e1)

	// Nested error.
	fmt.Println("\nNested error:")
	e2 := errors.E(errors.Op("Read"), path, user, errors.Other, e1)
	fmt.Println(e2)

	// Output:
	//
	// Simple error:
	// Get: jane@doe.com/file: I/O error: network unreachable
	//
	// Nested error:
	// Read: jane@doe.com/file, user joe@blow.com: I/O error:
	//	Get: network unreachable

}
