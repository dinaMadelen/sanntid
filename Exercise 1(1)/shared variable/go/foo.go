// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    //"time"
)

var i = 0

func incrementing(c chan int) {
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {   
        i += 1
    }

    c <- i
}

func decrementing(decrementdone chan int) {
    //TODO: decrement i 1000000 times
    for k := 0; k < 1000000; k++ {
        i -= 1
    }
    decrementdone <- 0
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)    

    decrementdone := make(chan int)
    c := make(chan int)

    // TODO: Spawn both functions as goroutines
    go decrementing(decrementdone)

    select {
    case <- decrementdone:
        go incrementing(c) 
    }

    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    //time.Sleep(500*time.Millisecond)

    i := <-c
    

    Println("The magic number is:", i)
}
