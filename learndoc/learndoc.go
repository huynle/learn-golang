// This package is dedicated to learn how to effective document code in golang. There are a few things we will try to figure out in this package.
// More doc here.
//
// 1. Document a package
//
// 2. Document functions
//
// 3. Document methods
//
// 4. Visiblity
//
// See Also
//
// This is how to create a Hello World:
// To document code, an extra space is needed
//  fmt.Printf("Hello, World")
// Or:
//  fmt.Printf("Hello, " + name)
//
package learndoc

import "fmt"

// This is the host
const Host = "example.com"

// This appears under the const.
const (
	// This causes Foo to happen.
	OptionFoo = 1

	// This causes Bar to happen.
	OptionBar = 2

	// Documented, but not visible.
	optionSecret = 3
)

// Apple have fields that describe the apple attributes
type Apple struct {
	// The color of the apple is important to keep track of
	Color string
	// Origin specifies the location for which the apple came from
	Origin string
	// Name of the apple type
	Name string
	// weight is not accessible and its documentation will not show up either
	weight float64
	// eaten mark
	Eaten bool
}

// UpdateOrigin changes the origin location of the apple. It will return true if the update was successful
//
// godoc notes: if the function is a method to a struct, then the documentation will group the methods together. These are orderred last in the document.
func (a *Apple) UpdateOrigin(origin string) bool {
	a.Name = origin
	return true
}

// Eat apple function, will set the state of apple to eaten.
//
// godoc notes: functions will automatically get orderred at the top of the documentation
func Eat(a Apple) {
	if a.Eaten != true {
		a.Eaten = true
	}
	fmt.Println("Already eaten")
}
