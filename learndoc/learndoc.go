// This package is dedicated to learn how to effectively document code in golang.
//
// Story
//
// Apple shop, have a bunch of apples. Each Apple is have attributes such as color, name, origin,
// weight, eaten, etc...
//
// This apple shop have access to all these attributes of an apple. But the customers can only see the name, and the color of the apple. While the shop owner can update the origin of the apple, the
//
// Extra Notes
//
// To document code, a space is needed after the "//" in the comment section
// This is how to create a Hello World:
//  fmt.Printf("Hello, World")
// Or:
//  fmt.Printf("Hello, " + name)
//
// Author: Huy Le
package learndoc

import "fmt"

// This is the host. Single constants will automatically get reorderred to the bottom of constants documentations.
const Host = "appleshop.com"

// This appears under the constant section.
const (
	// This causes Foo to happen.
	OptionFoo = 1

	// This causes Bar to happen.
	OptionBar = 2

	// Documented, but not visible.
	optionSecret = 3
)

// Apple have fields that describe the apple attributes
//
// Notes
//
// Methods are attached to their receiver type in the godoc, regardless of
// their physical location in the code.
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

// UpdateOrigin method have the same documentation style as UpdateColor
func (a *Apple) UpdateOrigin(origin string) bool {
	a.Origin = origin
	return true
}

// UpdateColor changes the origin location of the apple. It will return true if the update was successful
func (a Apple) UpdateColor(color string) bool {
	a.Color = color
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
