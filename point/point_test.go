/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:theme
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package point

import "testing"

// testpair documentation
type testpair struct {
	name   string
	expect string
}

// remember that var has to be used outside of a function to declare a variable
// BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
var tests = []testpair{
	{"foo", "foo"},
	{"bar", "bar"},
}

// TestFieldUpdate parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func TestFieldUpdate(t *testing.T) {
	for index, pair := range tests {
		ptr := Pointer{pair.name}
		if ptr.Update_ref_field() != pair.expect {
			t.Errorf("Expected %d, %s, got '%s' instead.", index, pair.name, ptr.name)
		}
	}
}

// this is a
func TestAllVarsArePointer(t *testing.T) {
	name := "foo"
	modify_field(&name)
}

//
// func TestStructUpdate(t *testing.T) {
// 	ptr := Pointer{"name"}
//
// 	if ptr.update_ref_struct() != "name" {
// 		t.Errorf("Expected name, got '%s' instead.", ptr.name)
// 	}
// }
