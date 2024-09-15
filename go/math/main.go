package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strings"
)

// 母分散σ**2
func getVariance(input []float64) float64 {
	if len(input) == 0 {
		return 0.0
	}

	N := len(input)
	AVE := getMean(input)

	var total float64
	for _, n := range input {
		total += math.Pow(n-AVE, 2)
	}

	return total * (1 / float64(N))
}

// 標本分散
func getSampleVariance(input []float64) float64 {
	if len(input) == 0 {
		return 0.0
	}

	N := len(input)
	AVE := getMean(input)

	var total float64
	for _, n := range input {
		total += math.Pow(float64(n)-AVE, 2)
	}

	return total * (1 / float64(N-1))
}

// 標準偏差σ
func getSigma(input []float64) float64 {
	return math.Sqrt(getVariance(input))
}

func getMean(input []float64) float64 {
	if len(input) == 0 {
		return 0.0
	}

	var total float64
	for _, v := range input {
		total += v
	}

	return float64(total) / float64((len(input)))
}

func getMedian(input []float64) float64 {
	if len(input) == 0 {
		return 0.0
	}

	cpInput := make([]float64, len(input))
	copy(cpInput, input)
	slices.Sort(cpInput)
	if len(cpInput)%2 == 1 {
		return cpInput[len(cpInput)/2]
	}

	return getMean([]float64{
		cpInput[(len(cpInput)/2)-1],
		cpInput[(len(cpInput) / 2)],
	})
}

func getScaled(inputs []float64) []float64 {
	ave := getMean(inputs)
	sigma := getSigma(inputs)

	result := make([]float64, len(inputs))
	for i := range inputs {
		result[i] = (float64(inputs[i]) - ave) / sigma
	}

	return result
}

func main() {
	{
		group := []float64{-2, -1, 0, 1, 2}
		fmt.Println("平均", getMean(group))
		fmt.Println("母分散", getVariance(group))
		fmt.Println("標準偏差:", getSigma(group))
	}
	{
		group := []float64{-4, -2, 0, 2, 4}
		fmt.Println("平均", getMean(group))
		fmt.Println("母分散", getVariance(group))
		fmt.Println("標準偏差:", getSigma(group))
	}

	for range 10 {
		values := make([]float64, 10)
		for i := 0; i < len(values); i++ {
			values[i] = rand.Float64()
		}

		scaled := getScaled(values)
		fmt.Println("scaled:", scaled)
		fmt.Printf("scaled mean: %.4f\n", getMean(scaled))
		fmt.Printf("scaled sigma %.4f\n:", getSigma(scaled))
		fmt.Println(strings.Repeat("*", 100))
	}

}
