package shape

import "fmt"

//Second Object Circle
type Square struct {
	// there is NO inheritance here! the object must be created the developer would have to make sure that the struct have all the functions defined by the interface
	name string
	Side float32
}

// this little snippet will make golang do its check to make sure tha Square have all its method implmented from the interface
func newSquare(side float32) Shape {
	f := &Square{Side: side}
	return f
}

// this is like the @staticmethod in python
// it implements the Shape object
func (s *Square) Area() string {
	return fmt.Sprintf("%f", s.Side*s.Side)
	// return s.Side * s.Side
}

// this is like the @staticmethod in python
// golang would throw error if this method was not implemented
func (s *Square) Name() string {
	return fmt.Sprintf("%f", s.Name)
	// return s.Side * s.Side
}
