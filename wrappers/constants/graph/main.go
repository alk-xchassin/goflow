// Code generated by lib-go/goflow DO NOT EDIT.

// +build !codeanalysis

package main

import (
	"fmt"
)

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func main() {
	var g Constants

	g = NewConstants()
	g.Run()
	fmt.Println("Test done.")

	fmt.Println("All tests done.")
}