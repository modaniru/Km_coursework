package main

import "math"

var integrals = []Integral{
	{
		name:          "1/x",
		a:             1,
		b:             2,
		correctAnswer: 0.693,
		integralFunction: func(x float64) float64 {
			return 1 / x
		},
	},
	{
		name:          "x",
		a:             0,
		b:             1,
		correctAnswer: 0.5,
		integralFunction: func(x float64) float64 {
			return x
		},
	},
	{
		name:          "cos^3(x)",
		a:             0,
		b:             math.Pi / 2,
		correctAnswer: 2. / 3.,
		integralFunction: func(x float64) float64 {
			return math.Pow(math.Cos(x), 3)
		},
	},
	{
		name:          "|1-x|",
		a:             0,
		b:             2,
		correctAnswer: 1,
		integralFunction: func(x float64) float64 {
			return math.Abs(1 - x)
		},
	},
	{
		name:          "x^(1/3)",
		a:             -1,
		b:             8,
		correctAnswer: 45. / 4.,
		integralFunction: func(x float64) float64 {
			return math.Cbrt(x)
		},
	},
	{
		name:          "x(3x^4 - 4x^2 + 1)",
		a:             -1,
		b:             2,
		correctAnswer: 18,
		integralFunction: func(x float64) float64 {
			return x * (3*math.Pow(x, 4) - 4*math.Pow(x, 2) + 1)
		},
	},
}

func Sum(values []float64) float64 {
	res := 0.
	for _, v := range values {
		res += v
	}

	return res
}

func InInterval(v, start, end float64) bool {
	return v >= start && v <= end
}
