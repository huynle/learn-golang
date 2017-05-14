// share the same package as the interface
// this is the reason why Shape is not shape.Shapres
package shape

import "fmt"

//Second Object Circle
type Circle struct {
	// there is NO inheritance here! the object must be created the developer would have to make sure that the struct have all the functions defined by the interface
	name   string
	Radius float32
}

func newCircle(radius float32) Shape {
	f := &Circle{Radius: radius}
	return f
}

// this is like the @staticmethod in python
// it implements the Shape object
func (c *Circle) Area() string {
	return fmt.Sprintf("%f", 3.14*c.Radius*c.Radius)
	// return s.Side * s.Side
}

// this is like the @staticmethod in python
// golang would throw error if this method was not implemented
func (c *Circle) Name() string {
	return fmt.Sprintf("%f", c.name)
	// return s.Side * s.Side
}
