package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Integral struct {
	name             string
	integralFunction func(x float64) float64
	generatedValues  []float64
	processedValues  []float64
	correctAnswer    float64
	evaluation       float64
	dispersion       float64
	a                float64
	b                float64
}

func (i *Integral) GenerateRandomValues(n int) []float64 {
	res := make([]float64, 0, n)
	for range n {
		res = append(res, rand.Float64()*(i.b-i.a)+i.a)
	}

	return res
}

func (i *Integral) CalculateByFunc() []float64 {
	res := make([]float64, 0)
	for _, x := range i.generatedValues {
		res = append(res, i.integralFunction(x))
	}

	return res
}

func (i *Integral) MonteCarlo(n int) {
	u := i.GenerateRandomValues(n)
	i.generatedValues = u

	x := i.CalculateByFunc()
	i.processedValues = x

	i.evaluation = (i.b - i.a) * (Sum(x) / float64(len(x)))
}

func (i *Integral) CalculateDisperse() float64 {
	res := 0.
	n := float64(len(i.processedValues))
	for _, xi := range i.processedValues {
		res += math.Pow(xi-i.evaluation, 2) / (n - 1)
	}

	i.dispersion = res
	return res
}

func (i *Integral) CalculateInterval() (float64, float64) {
	sqrtS := math.Sqrt(i.dispersion)
	z := 1.96

	start := i.evaluation - (z*sqrtS)/math.Sqrt(float64(len(i.processedValues)))
	end := i.evaluation + (z*sqrtS)/math.Sqrt(float64(len(i.processedValues)))

	return start, end
}

func main() {
	n := 10_000_000
	showTable := false

	for _, i := range integrals {
		i.MonteCarlo(n)
		i.CalculateDisperse()

		if showTable {
			fmt.Printf("\nТаблица для функции: %s\n", i.name)
			fmt.Printf("| %-3s | %-10s | %-15s |\n", "n", "xi", "f(xi)")
			fmt.Println("|------|------------|-----------------|")

			for j := 0; j < n; j++ {
				fmt.Printf("| %-3d | %-10.5f | %-15.5f |\n", j+1, i.generatedValues[j], i.processedValues[j])
			}
		}

		// Вывод основной информации по интегралу
		start, end := i.CalculateInterval()
		fmt.Println("\nРезультаты для функции:", i.name, "[", i.a, ",", i.b, "]")
		fmt.Printf("I (точное значение): %.3f\n", i.correctAnswer)
		fmt.Printf("W (оценка): %.5f\n", i.evaluation)
		fmt.Printf("|I - W|: %.5f\n", math.Abs(i.evaluation-i.correctAnswer))
		fmt.Printf("S^2 (дисперсия): %.5f\n", i.dispersion)
		fmt.Printf("Доверительный интервал: [%.5f : %.5f]\n", start, end)
		if InInterval(i.evaluation, start, end) {
			fmt.Println("Находится в доверительном интервале")
		} else {
			fmt.Println("Не находится в доверительном интервале")
		}

		fmt.Println()
	}
}
