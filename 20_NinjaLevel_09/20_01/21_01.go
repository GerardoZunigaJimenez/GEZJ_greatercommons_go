package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go goRuntine1()

	go goRuntine2()

	wg.Wait()
}

func goRuntine1() {
	for i := 0; i<10; i++ {
		fmt.Println("goRuntine1:", i)
	}

	wg.Done()
}


func goRuntine2() {
	for i := 0; i<10; i++ {
		fmt.Println("goRuntine2:", i)
	}

	wg.Done()
}