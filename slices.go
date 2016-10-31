// hello.go
package main

import ( 
	"fmt"
)

func main() {
	foo := []int{1,2,3,4,5,6}
	fmt.Println( foo[:2] )
	fmt.Println( foo[2:4] )
	fmt.Println( foo[4:] )
}
