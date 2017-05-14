package fruit

import "fmt"

const typical_halflife = 10.1

// const (
// 	// Typical halflife of an apple
// 	// should be able to overwrite if a specific type of apple rots faster
// 	halflife_grannysmith float64 = 10
// )

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
	// apples CAN have different halflife
	halflife float64
}

// sets up the apple package to have the correct halflife for fruit
func init() {
	fmt.Println("Setting up the Apple!")
	halflife = typical_halflife
}

// this little snippet will make golang do its check to make sure tha Square have all its method implmented from the interface
//
// when creating this function, the arguments does not have to be specified. This really is here to force the struct to implement all the methods of the struct.
func newApple() Fruit {
	apple := &Apple{}
	return apple
}

// When an apple rot, it rots at the rate of its halflife.
//
// In general, most apple rot at the same rate, expect for those grannysmiths!
func (a *Apple) GetDaysUntilRot() float64 {
	// TODO(huy): need to calculate days until the apple is rotted
	if a.halflife != 0 {
		return a.halflife
	} else {
		return halflife
	}
}

func (a *Apple) harvested() string {
	// TODO(huy): Need to get this going
	return "harvested date"
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
