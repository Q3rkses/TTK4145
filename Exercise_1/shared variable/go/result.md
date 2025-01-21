For the following GO code if we set GOMAXPROCS to 1 we get the same (bad) solution as in C, a single thread 
running at any given time and we essencially get rid of all concurrency.

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

--------------------------------------Concurrency solution:--------------------------------------

The second iteration of the code will make use of channels, in this case we use the select function to be able to recieve a message once ANY source is ready to send, and afterward we simply increase or decrease i inside of this select function. In this case both threads will intermingle however concurrency is kept in the way that only one increase or decrease can happen at one time to the variable i.