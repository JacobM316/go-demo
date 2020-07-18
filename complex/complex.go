package complex

import "fmt"
import "math"

type Complex struct {
	real float64
	imag float64
}

func (c *Complex) Real() float64 {
	return c.real
}

func (c *Complex) Imag() float64 {
	return c.imag
}

func (c *Complex) SetReal(r float64) {
	c.real = r
}

func (c *Complex) SetImag(i float64) {
	c.imag = i
}

func (c *Complex) Add(other Complex) Complex {
	return Complex{c.real + other.real, c.imag + other.imag}
}

func (c *Complex) Sub(other Complex) Complex {
	return Complex{c.real - other.real, c.imag - other.imag}
}

func (c Complex) String() string {
	var realSign, imagSign string
	if c.real < 0 {
		realSign = "-"
	} else {
		realSign = ""
	}

	if c.imag < 0 {
		imagSign = "-"
	} else {
		imagSign = "+"
	}
	return fmt.Sprintf("(%s%v%s%vi)", realSign, math.Abs(c.real), imagSign, math.Abs(c.imag))
}

func NewComplex(r float64, i float64) Complex {
	return Complex{r, i}
}
