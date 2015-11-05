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
		
