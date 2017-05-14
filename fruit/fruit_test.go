package fruit

import "testing"

type testpair struct {
	name     string
	origin   string
	color    string
	halflife float64
}

// generating test pairs
var tests = []testpair{
	{"Grannysmith", "Denver, CO", "blue", 1.1},
	{"Honeycrisp", "Portland, OR", "red", 0},
	{"Fuji", "Southern, CA", "yellow", 12.4},
}

// Ensure the fields of the Apple struct can be by key value pair.
func TestName(t *testing.T) {
	for _, pair := range tests {
		apple := Apple{Name: pair.name}
		t.Log("My apple name is: %s", apple.Name) // t.Log allows us to report log if error occurs
		if apple.Name != pair.name {
			t.Errorf("problem here")
		}
	}
}

//
func TestRot(t *testing.T) {
	for _, pair := range tests {
		apple := Apple{halflife: pair.halflife}
		// for the case the 0 is set for halflife
		if pair.halflife == 0 {
			if apple.GetDaysUntilRot() != typical_halflife {
				t.Errorf("%s halflife is not set, expecting typical %f halflife, but got %f", pair.name, typical_halflife, apple.GetDaysUntilRot())
			}
		} else if apple.GetDaysUntilRot() != pair.halflife {
			t.Errorf("%s is halflife is %f, while the expectation is %f", pair.name, apple.GetDaysUntilRot(), pair.halflife)
		}
	}
}
