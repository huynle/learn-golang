// This is the entry poin to the program.
// A few cool things can take place from here. The first thing to do is to read the comments.
// but let's try to read the comment from a served webpage.
// run `godoc -http=:6060` and go to `http://localhost:6060/pkg/github.com/huynle/learn-golang/`
package main

import (
	"fmt"
	"os"
)

// some var to track
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

// Emptyfunction is used to test the documentation of golang
func Emptyfunction() {
	fmt.Println("Not empty")
}

func main() {
	fmt.Println("Current Home Path is: %s", home)
	fmt.Println("Current User is: %s", user)
	fmt.Println("Current GOPATH is: %s", gopath)
	// testSquare1 := shape.Square{Side: 4}
	// fmt.Println("Area of a Square: " + testSquare1.Area())
	//
	// testCircle := shape.Circle{Radius: 4}
	// fmt.Println("Area of a Circle: " + testCircle.Area())
}
