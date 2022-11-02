package main

import (
	"fmt"
	"sync"
)

func g1() {
	fmt.Println("G1")
	wg.Done()
}

func g2() {
	fmt.Println("G2")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin")

	wg.Add(2)

	go g1()
	go g2()

	wg.Wait()

	fmt.Println("End")

}
