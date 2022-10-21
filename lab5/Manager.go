package main

import (
	"fmt"
	"math/rand"
)

var LIMIT = 5

type Manager struct {
	array      []int
	currentSum chan int
}

func NewManager(size int) *Manager {
	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(int(LIMIT))
	}

	newManager := &Manager{array, make(chan int, 1)}
	newManager.evaluateSum()

	return newManager
}

func (thread *Manager) getCurrentSum() int {
	sum := <-thread.currentSum
	thread.currentSum <- sum
	return sum
}

func (thread *Manager) evaluateSum() {
	sum := 0
	for _, currentItem := range thread.array {
		sum += currentItem
	}

	if len(thread.currentSum) == 1 {
		<-thread.currentSum
	}
	thread.currentSum <- sum
}

func (thread *Manager) arrayModification() {
	index := rand.Intn(len(thread.array))

	if rand.Intn(2) == 0 {
		thread.array[index] = (thread.array[index] - 1) % int(LIMIT)
	} else {
		thread.array[index] = (thread.array[index] + 1) % int(LIMIT)
	}

	thread.evaluateSum()
}

func (thread *Manager) print() {
	sum := thread.getCurrentSum()
	fmt.Printf("Elements: %v\nSum: %d\n", thread.array, sum)
}
