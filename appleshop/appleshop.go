// This package is dedicated to learn how to effectively document code in golang.
//
// The Apple Story
//
// Apple shop, have a bunch of apples. Each Apple is have attributes such as color, name, origin,
// weight, eaten, etc...
//
// This apple shop has a shop owner and some customers. The shopowner can decide to place his store online, or just in a local shop. Regardless, he has the option to place everything on sale, just in case the apples are about to expire.
//
// When the customers buy these apples, they only know the color of the apple, size and weight of the apple. When they make a decision to purchase,they can either buy it online, or in the local shop. Once bought, they can do whatever they want to the apple, i.e. eat it, store it, throw it away, etc.
//
// Notes
//
// This story line will help us create a package that makes sense and is coherrent to the story.
//
// Author: Huy Le
package appleshop

import "fmt"

// This is the base URL for the appleshop online.
const Host = "appleshop.com"

// A series of constants are created here.
//
// Since constants are defined during compile time. These constants are used to set up the appleshop.
const (
	// if set to true, this constant will set up the onlinestore
	onlineStore = false

	// When the owner want to have a flash sale and clear the apply inventory.
	FlashSale = false
)

// Eat apple function, will set the state of apple to eaten.
//
// godoc notes: functions will automatically get orderred at the top of the documentation
func Eat(a Apple) {
	if a.Eaten != true {
		a.Eaten = true
	}
	fmt.Println("Already eaten")
}
