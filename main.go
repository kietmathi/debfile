package main

import (
	"cloud.google.com/go/storage"
	"fmt"
)

func main() {
	fmt.Println(storage.ScopeReadWrite)
}
