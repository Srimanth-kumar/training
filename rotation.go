package main

import (
	"log"
)

func main() {

	narr := []string{"one", "two", "three", "four", "five", "six"}

	log.Println(narr)

	rotate(narr, 2)

	log.Println(narr)

}

func rotate(a []string, idx int) {

	//cre a new array that can hold all of the elements ( size=len of a)
	temp := make([]string, len(a))
	//copy elements one by one
	copy(temp, a)

	// loop from idx to length and then from 0 to idx
	for i := idx; i < len(a)+idx; i++ {
		temp[i-idx] = a[i%len(a)]
	}

	//update the original array
	copy(a, temp)
}
