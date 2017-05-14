// this package definition let golang know the files that are associated with the package
package shape

// this sets the interface for the shape object
type Shape interface {
	// defining must have functions for structs that is of Shape
	Area() string
	Name() string
	// golang would throw error if a method is method of the interface is defined here and it is not implemented
}
