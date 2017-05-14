// point package is about exploring the pointers in golang
package point

import "fmt"

// more doc on functions
func Modify_field(name *string) {
	*name = "field"
}

// another one on function
func Modify_struct(ptr *Pointer) {
	ptr.name = "struct"
}

// Some doc on Pointer
type Pointer struct {
	name string
}

// init is the start
func init() {
	fmt.Println("start point package")
}

// only the Caps are documented
func (ptr *Pointer) Update_ref_field() string {
	Modify_field(&ptr.name)
	return ptr.name
}

// some other doc here
func (ptr *Pointer) Update_ref_struct() string {
	Modify_struct(ptr)
	return ptr.name
}
