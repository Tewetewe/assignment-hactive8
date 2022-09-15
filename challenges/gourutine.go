package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	//uncomment if you want to try unrace condition

	// var mu sync.Mutex

	listWord1 := []string{"bisa1", "bisa2", "bisa3", "bisa4"}
	listWord2 := []string{"coba1", "coba2", "coba3", "coba4"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)

		//unrace condition  (uncomment for unrace condition)

		// go PrintInterface1(i, listWord1, &wg, &mu)
		// go PrintInterface2(i, listWord2, &wg, &mu)

		//random (uncomment for random queue)

		go PrintInterface3(i, listWord1, &wg)
		go PrintInterface4(i, listWord2, &wg)
	}

	wg.Wait()

}

func PrintInterface1(index int, listWord interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()

	fmt.Println(listWord, index)

	mu.Unlock()

	wg.Done()
}

func PrintInterface2(index int, listWord interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()

	fmt.Println(listWord, index)

	mu.Unlock()

	wg.Done()
}

func PrintInterface3(index int, listWord interface{}, wg *sync.WaitGroup) {

	fmt.Println(listWord, index)

	wg.Done()
}

func PrintInterface4(index int, listWord interface{}, wg *sync.WaitGroup) {

	fmt.Println(listWord, index)

	wg.Done()
}
