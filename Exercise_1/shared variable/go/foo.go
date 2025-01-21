// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing(c_incr chan int) {
    for j := 0; j < 1000000; j++ {
        c_incr <- 1
    }
    close(c_incr)
}

func decrementing(c_decr chan int) {
    for j := 0; j < 1000001; j++ {
        c_decr <- -1
    }
    close(c_decr)
}

func main() {
    runtime.GOMAXPROCS(2) // Controls how many threads can run simultaneously, if we set to 1 then the output will be 0

    c_incr := make(chan int)
    c_decr := make(chan int)

    go incrementing(c_incr)
    go decrementing(c_decr)

    for (c_decr != nil || c_incr != nil) {
        select {
        
        case increment, ok := <-c_incr:
            i += increment

            if !ok {
                c_incr = nil
            }

        case decrement, ok := <-c_decr:
            i += decrement

            if !ok {
                c_decr = nil
            }   
        }
    }

    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
