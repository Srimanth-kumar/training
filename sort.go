package main

import (
	"log"
	"time"
)

func main() {

	narray := []int{10, 20, 30, 40, 15, 2, 3, 8, 20}

	dump("before Sort", narray)

	time1 := time.Now()
	sortArray(narray)
	log.Println("Time taken:", time.Since(time1))

	dump("aftersort", narray)

}

func dump(s string, a []int) {
	log.Print(s, a)

}

func sortArray(a []int) {

	//TODO: Write sorting logic of a

	//time.Sleep(1 * time.Second)
}

//Excercises :
// 1. use basic sort algorithm
// 2. Try with larger array ( size 20)
// 3. Try with Generating 1000 random numbers into an array and sort
