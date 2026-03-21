// Possible arithmetic operations for intigers
package main

import "fmt"

func main() {

	var x int16 = 170
	var y int16 = 83
	//Addition
	fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
	//Subtraction
	fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
	//Multiplication
	fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
	//Division
	fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)

	//Modulus
	fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)

}
