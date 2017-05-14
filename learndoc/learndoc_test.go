package learndoc

import (
	"fmt"
	"testing"
)

type testpair struct {
	name   string
	origin string
}

var tests = []testpair{
	{"Granny Smith", "Denver, CO"},
	{"Honeycrisp", "Portland, OR"},
}

// Test eating apple
func TestEatingApple(t *testing.T) {
	for _, pair := range tests {
		apple := Apple{Name: pair.name}
		fmt.Println(apple.Name)
	}
}

// This is a struc level example
func ExampleApple() {
	your_apple := Apple{Name: "Example Name"}
	fmt.Println("Here's your apple! : %s", your_apple)
}

// This is a function level example
func ExampleEat() {
	my_apple := Apple{Name: "Example Name"}
	fmt.Println("Is your apple eaten yet? ", my_apple.Eaten)
	Eat(my_apple)
}

// This is a package-level example with an assertsion to test the example code and keeping it consistent
func Example() {
	my_apple := Apple{Name: "Example Name"}
	fmt.Println(my_apple.Name)
	// Output: Example Name
}
