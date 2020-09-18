package main

import (
    "sync"
    "time"
)

func main() {
    ch := make(chan int)

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go runChan(wg, ch)

    for j := 0; j < 9; j++ {
        ch <- j
    }
    close(ch)
    wg.Wait()

    time.Sleep(100 * time.Millisecond)
}

func runChan(wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    for { //i := 0; i < 10; i++
        select {
        case i, ok := <-ch:
            if !ok {
                goto end
            }
            println(i)

        default:
            println("default")
        }
    }
end:
    println("end")
}
