package main

import (
	"fmt"
	"math"
	"strings"
)

// 母分散σ**2
func getVariance(input []int) float64 {
	if len(input) == 0 {
		return 0.0
	}

	N := len(input)
	AVE := getAverage(input)

	var total float64
	for _, n := range input {
		total += math.Pow(float64(n)-AVE, 2)
	}

	return total * (1 / float64(N))
}

// 標本分散
func getSampleVariance(input []int) float64 {
	if len(input) == 0 {
		return 0.0
	}

	N := len(input)
	AVE := getAverage(input)

	var total float64
	for _, n := range input {
		total += math.Pow(float64(n)-AVE, 2)
	}

	return total * (1 / float64(N-1))
}

// 標準偏差σ
func getSigma(input []int) float64 {
	return math.Sqrt(getVariance(input))
}

func getAverage(input []int) float64 {
	if len(input) == 0 {
		return 0.0
	}

	var total int
	for _, v := range input {
		total += v
	}

	return float64(total) / float64((len(input)))
}

func main() {
	{
		group := []int{-2, -1, 0, 1, 2}
		fmt.Println("平均", getAverage(group))
		fmt.Println("母分散", getVariance(group))
		fmt.Println("標準偏差:", getSigma(group))
	}
	fmt.Println(strings.Repeat("*", 100))
	{
		group := []int{-4, -2, 0, 2, 4}
		fmt.Println("平均", getAverage(group))
		fmt.Println("母分散", getVariance(group))
		fmt.Println("標準偏差:", getSigma(group))
	}
}
