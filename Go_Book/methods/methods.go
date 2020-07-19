package main

import (
	"fmt"
	"math"
)

// VV is not a W :-P
type VV struct {
	// X this, Y that.
	X, Y float64
}

// Abs is the sqrt of X^2+Y^2. A method is just a function with a
// receiver argument. This way can call struct.something. If Abs was
// abs (vv VV, h string) then would call it like abs(vv, "extra_arg").
func (vv VV) Abs(h string) float64 {
	fmt.Println(h)
	return math.Sqrt(vv.X*vv.X + vv.Y*vv.Y)
}

// MyFloat will be an alias so that we can declare methods on primitive
// types such as float64.
type MyFloat float64

// Abs is just a simple method for a MyFloat receiver but could be more
// complex stuff. Type has to be MyFloat.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// can call method like VV{X: 3, Y: 4}.Abs()
	// The first argument of the method will be the struct.
	// Then any other number of arguments will be passed.
	vv := VV{X: 3, Y: 4}
	fmt.Println(vv.Abs("extra_arg"))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
