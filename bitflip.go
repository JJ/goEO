// hello.go
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
	for length <= 32768 {
		vec := []bool{}
		i := 0
		start := time.Now()
		for i < length {
			this_bool := true
			if ( rand.Float32() < 0.5 ) {
				this_bool = false
			}
			vec = append( vec, this_bool )
			i++
		}
		new_vec := vec
		iter := 0
		for iter < iterations {
			new_vec = mutate(new_vec)
			iter++
		}
		fmt.Println("Go-BitVector,",length,", ", time.Since(start).Seconds())

		// Crossover
		vec2 := []bool{}
		i = 0
		start = time.Now()
		for i < length {
			this_bool := true
			if ( rand.Float32() < 0.5 ) {
				this_bool = false
			}
			vec2 = append( vec2, this_bool )
			i++
		}
		new_vec_2 := vec2
		iter = 0
		for iter < iterations {
			new_vec, new_vec_2 = crossover(new_vec, new_vec_2)
			iter++
		}
		fmt.Println("Go-BitVector,",length,", ", time.Since(start).Seconds())
		length = length*2
	} 
}

func mutate( array []bool ) []bool {
	point_of_mutation  := rand.Intn(len(array))
	result := []bool{}
	if ( point_of_mutation > 0 ) {
		result = array[:point_of_mutation]
	}
	if ( array[point_of_mutation] ) {
		result = append(result,false)
	} else {
		result = append(result,true)
	}
	result = append(result,array[point_of_mutation+1:]...)
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
