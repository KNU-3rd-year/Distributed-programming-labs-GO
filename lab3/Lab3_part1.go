package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("enter the capacity of the jar: ")
	fmt.Scan(&jarCapacity)

	fmt.Printf("enter the size of the beehive: ")
	var beehiveSize int
	fmt.Scan(&beehiveSize)

	ch := make(chan int, jarCapacity)

	for counter.getCounter() == 0 {
		for i := 0; i < beehiveSize; i++ {
			if counter.getCounter() == 1 {
				break
			}
			fmt.Printf("Bee %v is heading to the jar\n", i+1)
			waitGroup.Add(1)
			go fill(ch)
		}
	}

	waitGroup.Wait()
	fmt.Println("Bear has eaten everything")
}

var jarCapacity int
var counter = new(Counter)
var waitGroup sync.WaitGroup

type Counter struct {
	sync.Mutex
	count int
}

func fill(ch chan int) {
	defer waitGroup.Done()
	if counter.getCounter() == 1 {
		return
	}
	select {
	case ch <- len(ch):
		fmt.Printf("Filling the jar (%v/%v)\n", len(ch), jarCapacity)
	default:
		counter.incrementCounter()
		awakeAndEat(ch)
	}
}

func awakeAndEat(ch chan int) {
	close(ch)
	fmt.Printf("Jug is full %v\n", len(ch))
	fmt.Println("Bear awake")
	fmt.Printf("Bear is eating honey...\n")
}

func (c *Counter) incrementCounter() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) getCounter() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}
