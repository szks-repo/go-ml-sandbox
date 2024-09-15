package main

import (
	"fmt"
	"math/rand"
)

type Product struct {
	ID   int
	Name string
}

type Arm struct {
	Product Product
	Tries   int
	Rewards int
	CTR     float64
}

type MultiArmedBandit struct {
	Arms       []Arm
	Epsilon    float64
	TotalTries int
}

func NewMultiArmedBandit(products []Product, epsilon float64) *MultiArmedBandit {
	arms := make([]Arm, len(products))
	for i, product := range products {
		arms[i] = Arm{Product: product}
	}
	return &MultiArmedBandit{
		Arms:    arms,
		Epsilon: epsilon,
	}
}

func (mab *MultiArmedBandit) SelectArm() *Arm {
	if mab.TotalTries == 0 || rand.Float64() < mab.Epsilon {
		// Exploration: 乱数で腕を選択
		return &mab.Arms[rand.Intn(len(mab.Arms))]
	}

	// Exploitation: 最高のCTRを持つ腕を選択
	var bestArmPos int
	var maxCTR float64
	for i := range mab.Arms {
		if mab.Arms[i].CTR > maxCTR {
			maxCTR = mab.Arms[i].CTR
			bestArmPos = i
		}
	}

	return &mab.Arms[bestArmPos]
}

func (mab *MultiArmedBandit) Update(arm *Arm, reward int) {
	arm.Tries++
	arm.Rewards += reward
	arm.CTR = float64(arm.Rewards) / float64(arm.Tries)
	mab.TotalTries++
}

func main() {
	products := []Product{
		{ID: 1, Name: "商品A"},
		{ID: 2, Name: "商品B"},
		{ID: 3, Name: "商品C"},
		{ID: 4, Name: "商品D"},
		{ID: 5, Name: "商品E"},
		{ID: 6, Name: "商品F"},
		{ID: 7, Name: "商品G"},
		{ID: 8, Name: "商品H"},
		{ID: 9, Name: "商品I"},
		{ID: 10, Name: "商品J"},
	}

	mab := NewMultiArmedBandit(products, 0.1) // イプシロン = 0.1

	// シミュレーション: 10000回の表示
	for i := 0; i < 10000; i++ {
		arm := mab.SelectArm()

		reward := 0
		if rand.Float64() < 0.1+float64(arm.Product.ID)*0.01 {
			reward = 1
		}

		mab.Update(arm, reward)
	}

	fmt.Println("シミュレーション結果:")
	for _, arm := range mab.Arms {
		fmt.Printf("商品ID: %d, 名前: %s, 試行回数: %d, CTR: %.4f\n",
			arm.Product.ID, arm.Product.Name, arm.Tries, arm.CTR)
	}

	tests := make(map[int]int, 10)
	for _, arm := range mab.Arms {
		tests[arm.Product.ID] = 0
	}
	for range 1000 {
		bestArm := mab.SelectArm()
		tests[bestArm.Product.ID]++
	}

	fmt.Println(tests)
}
