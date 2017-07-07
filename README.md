Genetic Algorithm that solves the knapsack problem. Optimal solution is not guaranteed. It depends on the settings you used and some luck ;)

## Installation

```bash
$ git clone https://github.com/azaryc2s/ga_knapsack
$ cd ga_knapsack
$ go build
```

## Usage

This software reads the data from the stdin until EOF and prints the output to the stdout. So you either have to use echo and pipe the output to it or simply pass it a file like so:

```bash
$ ./ga_knapsack < in
```
Normally you'd like to store the output somewhere so you'd use
```bash
$ ./ga_knapsack < in > out
```

## Input

The input has to be in the json format. The following snippet contains an overview of the parameters that you can specify. Values that do not have a default value have to be set by you. 

```json
{
	"Profits": [int-array],
	"Weights": [float-array],
	"Capacity": [int],	
	"Population": [int], //default is 20
	"MutRate": [float], //default is 0.1
	"Generations": [int] //default is 5000
}
```

## Output

The output is also in json format. It contains 3 values. 'Knapsack' is a bit array (actually an int-array only containing 0 or 1) which sets, which item has made it to the knapsack (in the same order as specified in the input). 'Weight'

```json
{
	"Knapsack": [bit-array], //Array of 0 or 1 specifing which item should be included in the knapsack
	"Weight": int, //Sum of the item weights included in the knapsack
	"Profit": int //Sum of the item profits included in the knapsack
}
```
