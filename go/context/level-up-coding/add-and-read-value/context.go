package main

import (
	"context"
	"fmt"
	"log"
)

// use custom type to avoid collision.
type keyType string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, keyType("key"), "test-value")
	val, ok := ctx.Value(keyType("key")).(string)
	if !ok {
		log.Fatal("Fail to get the key: type assertion error")
	}
	fmt.Println(val)
}
