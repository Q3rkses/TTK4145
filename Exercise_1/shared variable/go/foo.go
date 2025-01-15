// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    for j := 0; j < 1000000; j++ {
        i++
    }
}

func decrementing() {
    for j := 0; j < 1000000; j++ {
        i--
    }
}

func main() {
    runtime.GOMAXPROCS(2) // Controls how many threads can run simultaneously, if we set to 1 then the output will be 0
	
    go incrementing()
    go decrementing()

    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
