package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	var chs = make(chan int)
	result := 0

	waitGroup := sync.WaitGroup{}
	for j := 0; j < 3; j++ {
		waitGroup.Add(1)
		go func(group *sync.WaitGroup) {
			for v := range chs {
				result += v
			}
			waitGroup.Done()
		}(&waitGroup)
	}

	for i := 0; i < 100; i++ {
		chs <- i
	}

	close(chs)
	waitGroup.Wait()

	log.Println(fmt.Sprintf("results: %v", result))
}
