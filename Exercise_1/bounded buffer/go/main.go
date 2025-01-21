
package main

import "fmt"
import "time"


func producer(buff chan int){

    for i := 0; i < 10; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        buff <- i
    }

    close(buff)
}

func consumer(buff chan int){

    time.Sleep(500 * time.Millisecond)
    for {
        i, ok := <-buff

        if ok {
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(500 * time.Millisecond)
        }
    }

}


func main(){
    buff := make(chan int, 5)
    
    go consumer(buff)
    go producer(buff)
    
    select {}
}