package complexnumbers

import "math"

func pow2(n float64) float64 {
	return n * n
}

type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imaginary + n2.imaginary}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imaginary - n2.imaginary}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		n1.real*n2.real - n1.imaginary*n2.imaginary,
		n1.imaginary*n2.real + n1.real*n2.imaginary,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		n.real * factor,
		n.imaginary * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	divisor := pow2(n2.real) + pow2(n2.imaginary)
	return Number{
		(n1.real*n2.real + n1.imaginary*n2.imaginary) / divisor,
		(n1.imaginary*n2.real - n1.real*n2.imaginary) / divisor,
	}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.imaginary}
}

func (n Number) Abs() float64 {
	return math.Sqrt(pow2(n.real) + pow2(n.imaginary))
}

func (n Number) Exp() Number {
	factor := math.Exp(n.real)
	return Number{
		factor * math.Cos(n.imaginary),
		factor * math.Sin(n.imaginary),
	}
}
