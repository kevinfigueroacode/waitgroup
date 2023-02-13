//filename: main.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	func() {
		defer wg.Done()
fmt.Println("hola")
	}()
	func() {
		defer wg.Done()
fmt.Println("ni hao")
	}()
	func() {
		defer wg.Done()
fmt.Println("hello")
	}()
	wg.Wait()//wait until wg becomes 0
}