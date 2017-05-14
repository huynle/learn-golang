package shape_test

import "testing"

// func ExampleReverse() {
// 	fmt.Println(stringutil.Reverse("hello"))
// 	// Output: olleh
// }

type testpair struct {
	value float64
	area  float64
}

var tests = []testpair{
	{1, 1},
}

func TestFirst(t *testing.T) {
	t.Fail()
	// for _, pair := range tests {
	// v := pair.value
	// 	shape = Square(v)
	// 	if shape.Area() != pair.area {
	// 		t.Error(
	// 			"For", pair.values,
	// 			"expected", pair.average,
	// 			"got", v,
	// 		)
	// 	}
	// }
}

//
// func TestScan(t *testing.T) {
// 	for _, r := range readers {
// 		t.Run(r.name, func(t *testing.T) {
// 			testScan(t, r.f, Fscan)
// 		})
// 	}
// }
