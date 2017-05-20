package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Input struct {
	Weights     []float64
	Profits     []float64
	Capacity    float64
	MutRate     float64
	Population  int
	Generations int
}

type Output struct {
	Knapsack []int
	Weight   float64
	Profit   float64
}

var (
	input Input
)

func main() {
	//Default values
	input.Population = 20
	input.Generations = 5000
	input.MutRate = 0.1
//		input.Profits = []float64{150, 35, 200, 160, 60, 45, 60, 40, 30, 10, 70, 30, 15, 10, 40, 70, 75, 80, 20, 12, 50, 10}
//		input.Weights = []float64{9, 13, 153, 50, 15, 68, 27, 39, 23, 52, 11, 32, 24, 48, 73, 42, 43, 22, 7, 18, 4, 30}
//		input.Capacity = 400
	readInput()

	rand.Seed(time.Now().UTC().UnixNano())
	knapsacks := make([][]bool, input.Population)
	vals := make([]float64, input.Population)
	weights := make([]float64, input.Population)

	//generate random knapsacks according to the given population size
	for i := 0; i < input.Population; i++ {
		knapsacks[i] = randomKnapsack(len(input.Weights))
	}
	sortKnapsacks(knapsacks, vals, weights)
	for i := 0; i < input.Generations; i++ {
		for j := 0; j < input.Population; j++ {
			vals[j], weights[j] = evalKnapsackVal(knapsacks[j])
		}
		sortKnapsacks(knapsacks, vals, weights)
		for j := 0; j < (input.Population / 2); j++ {
			a := input.Population - 1 - j
			b := j + (input.Population / 2)
			knapsacks[j] = crossover(knapsacks[a], knapsacks[b])
			mutate(knapsacks[j])
		}
	}
	binKs := boolToBin(knapsacks[len(knapsacks)-1])
	res := &Output{Knapsack: binKs, Weight: weights[len(weights)-1], Profit: vals[len(vals)-1]}
	output, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}

func readInput() {
	data, err := ioutil.ReadAll(os.Stdin) // Read everything from standard input
	if err != nil {                       // and store the input in a string.
		fmt.Fprintf(os.Stderr, "%v", err) // If something get wrong:
		os.Exit(1)                        // exit with status code <> 0.
	}
	err = json.Unmarshal(data, &input) // Interprete the input as JSON and fill
	if err != nil {                    // the corresponding Go struct.
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func randomKnapsack(n int) []bool {
	k := 10000
	knapsack := make([]bool, n)
	for i := 0; i < n; i++ {
		knapsack[i] = rand.Intn(k) < (k / 2)
	}
	return knapsack
}

func evalKnapsackVal(knapsack []bool) (float64, float64) {
	sum := 0.0
	weightSum := 0.0
	for i := 0; i < len(knapsack); i++ {
		if knapsack[i] {
			sum = sum + input.Profits[i]
			weightSum = weightSum + input.Weights[i]
		}
	}
	if weightSum > input.Capacity {
		return 0.0, weightSum
	}
	return sum, weightSum
}

func (s SortSlice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
	(s.Knapsacks)[i], (s.Knapsacks)[j] = (s.Knapsacks)[j], (s.Knapsacks)[i]
	(s.Weights)[i], (s.Weights)[j] = (s.Weights)[j], (s.Weights)[i]
}

func sortKnapsacks(knapsacks [][]bool, vals []float64, weights []float64) {
	slice := &SortSlice{Float64Slice: vals, Knapsacks: knapsacks, Weights: weights, idx: make([]int, len(vals))}
	for i := range slice.idx {
		slice.idx[i] = i
	}
	sort.Sort(slice)
}

type SortSlice struct {
	sort.Float64Slice
	Knapsacks [][]bool
	Weights   []float64
	idx       []int
}

func crossover(kn1 []bool, kn2 []bool) []bool {
	r := 1 + rand.Intn(len(kn1)-2)
	offspring := make([]bool, len(kn1))
	for i := 0; i < r; i++ {
		offspring[i] = kn1[i]
	}
	for i := r; i < len(kn1); i++ {
		offspring[i] = kn2[i]
	}
	return offspring
}

func mutate(ks []bool) {
	k := 1000000
	for i := 1; i < len(ks); i++ {
		b := rand.Intn(k) <= int(float64(k)*input.MutRate)
		if b {
			ks[i] = !ks[i]
		}
	}
}

func boolToBin(ks []bool) []int {
	result := make([]int, len(ks))
	for i := 0; i < len(ks); i++ {
		if ks[i] {
			result[i] = 1
		}
	}
	return result
}
