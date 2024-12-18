package main

import (
	"fmt"
	"sync"
)

// func hello(wg *sync.WaitGroup) {
// 	fmt.Println("Hello, world!")
// 	wg.Done()
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go hello(&wg)
// 	wg.Wait()
// }

// func add(sum *int, n int, wg *sync.WaitGroup) {
// 	for i := 0; i < n; i++ {
// 	*sum += 1
// 	}
// 	wg.Done()
// }


// func main() {
// 	sum := 0
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 100; i++ {
// 	wg.Add(1)
// 	go add(&sum, 100000, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(sum)
// }

func add(sum *int, n int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	for i := 0; i < n; i++ {
	mtx.Lock()
	*sum += 1
	mtx.Unlock()
	}
	wg.Done()
}

   func main() {
	sum := 0
	var mtx sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
	wg.Add(1)
	go add(&sum, 100000, &wg, &mtx)
	}
	wg.Wait()
	fmt.Println(sum)
}