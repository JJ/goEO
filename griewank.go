// benchmark.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	iterations := 1000000
	dimensions := [7]int{2,4,6,8,10,20,100};

	// Onemax
	for i,_ := range dimensions {
		start := time.Now()
		for j := 0; j < iterations; j ++  {
			vec :=  random_chromosome(dimensions[i])
			griewank( vec )
//			fmt.Println( result )
		}
		fmt.Println("Go,",dimensions[i],", ", time.Since(start).Seconds())
	}

}

func random_chromosome( length int ) []float64 {
	chromosome := make([]float64,length)
	for i:= 0; i  < length; i ++  {
		chromosome[i] = rand.Float64()*1200.0-600.0
	}
	return chromosome
}

func griewank( vec []float64 ) float64 {
	partial_prod := 1.0
	partial_sum := 0.0
	for i,value := range vec {
		good_i := float64(i)+1
		partial_prod = partial_prod * math.Cos(value/math.Sqrt(float64(good_i)))
		partial_sum = partial_sum + value*value/4000.0
	}
	return ( partial_sum - partial_prod +1 )
}
