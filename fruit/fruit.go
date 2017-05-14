// The fruit package defines and interface for all the fruits in the world!
package fruit

// parameter to control how long the apple can stay fresh for
var halflife float64

// Fruit an interface that defines all the methods for a fruit.
type Fruit interface {
	// fruit can rot, and this method ensures that the fruit is rotted based on the type of fruit.
	GetDaysUntilRot() float64
	// harvest date
	harvested() string
}
