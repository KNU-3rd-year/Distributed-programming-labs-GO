package main

import (
	"fmt"
	"math/rand"
	"time"
)

func compute(array []int, lo int, hi int, c chan int) {
	if lo == hi {
		c <- array[lo]
		return
	}
	c1 := make(chan int)
	c2 := make(chan int)
	go compute(array, lo, (lo+hi)/2, c1)
	go compute(array, (lo+hi)/2+1, hi, c2)
	energy1 := <-c1
	energy2 := <-c2
	if energy1 > energy2 {
		c <- energy1
	} else {
		c <- energy2
	}
}

func start(array []int) int {
	c := make(chan int)
	go compute(array, 0, len(array)-1, c)
	return <-c
}

func main() {
	rand.Seed(time.Now().Unix())
	var numberOfPairs int
	fmt.Scan(&numberOfPairs)

	var qiEnergy []int
	for i := 0; i < numberOfPairs*2; i++ {
		qiEnergy = append(qiEnergy, rand.Int()%1000)
	}

	for i := 0; i < numberOfPairs*2; i++ {
		fmt.Print(qiEnergy[i])
		fmt.Print(" ")
	}
	fmt.Println(" ")

	result := start(qiEnergy)
	fmt.Println(result)
}
