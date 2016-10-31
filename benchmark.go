// benchmark.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	iterations := 100000
	length := 16

	// Onemax
	fmt.Println(" OneMax " )
	for length <= 32768 {
		start := time.Now()
		iter := 0
		for iter < iterations {
			vec :=  random_chromosome(length)
			onemax( vec )
			iter++
		}
		fmt.Println("Go-BitVector,",length,", ", time.Since(start).Seconds())
		length = length*2

	}

	// Mutation
	fmt.Print(" Mutation \n" )
	length = 16
	for length <= 32768 {
		vec := random_chromosome(length)
		new_vec := vec
		start := time.Now()
		iter := 0
		for iter < iterations {
			new_vec = mutate(new_vec)
			iter++
		}
		fmt.Println("Go-BitVector,",length,", ", time.Since(start).Seconds())
		length = length*2

	}

	// Crossover
	fmt.Print(" Crossover \n" )
	length = 16
	for length <= 32768 {
		vec := random_chromosome(length)
		new_vec := vec
		vec2 := random_chromosome(length)
		new_vec_2 := vec2
		start := time.Now()
		iter := 0
		for iter < iterations {
			new_vec, new_vec_2 = crossover(new_vec, new_vec_2)
			iter++
		}
		fmt.Println("Go-BitVector,",length,", ", time.Since(start).Seconds())
		length = length*2
	}
}

func random_chromosome( length int ) []bool {
	chromosome := make([]bool,length)
	for i:= 0; i  < length; i ++  {
		if ( rand.Float32() < 0.5 ) {
			chromosome[i] = true
		}
	}
	return chromosome
}

func mutate( array []bool ) []bool {
	point_of_mutation  := rand.Intn(len(array))
	result := array
	if array[point_of_mutation] {
		result[point_of_mutation] = false
	}else {
		result[point_of_mutation] = true
	}
	return result
}


func crossover( mate1 []bool, mate2 []bool ) ([]bool,[]bool) {
	point_of_xover_1  := rand.Intn(len(mate1))
	point_of_xover_2  := rand.Intn(len(mate1))
	if ( point_of_xover_1 > point_of_xover_2) {
		point_of_xover_1, point_of_xover_2 = point_of_xover_2, point_of_xover_1
	}
	result1 := []bool{}
	result2 := []bool{}
	if ( point_of_xover_1 > 0 ) {
		result1 = mate1[:point_of_xover_1]
		result2 = mate2[:point_of_xover_1]
	}
	result1 = append(result1,mate2[point_of_xover_1:point_of_xover_2]...)
	result2 = append(result2,mate1[point_of_xover_1:point_of_xover_2]...)

	result1 = append(result1,mate1[point_of_xover_2:]...)
	result2 = append(result2,mate2[point_of_xover_2:]...)
	return result1, result2
}

func onemax ( mate []bool ) int {
	ones := 0
	for _,element := range mate {
		if ( element ) {
			ones++
		}
	}
	return ones
}
