// Sample program to show how you can't always get the
// address of a value.
package main

import "fmt"

type p interface {
	pretty()
}

// duration is a type with a base type of int.
type duration struct {
	d int
}

// format pretty-prints the duration value.
func (d *duration) pretty() {
	fmt.Printf("Duration: %d", d.d)
}

// main is the entry point for the application.
func main() {
	d := duration{42}
	d.pretty()

	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}
