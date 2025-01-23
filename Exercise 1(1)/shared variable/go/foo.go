// Use `go run foo.go` to run your program

package main

import (
    "fmt"
    "runtime"
    //"time"
)

var i = 0

func incrementing( c chan int, done chan bool) {
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {   
        val :=<-c
        val += 1
        c <- val
    }

    done <- true
    
}

func decrementing(c chan int, done chan bool) {
    //TODO: decrement i 1000000 times
    for k := 0; k < 1000000; k++ {
        val :=<-c
        val -= 1
        c <- val
    }
    done <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)    

    c := make(chan int, 1)
    done := make(chan bool)
    c <- 0

    // TODO: Spawn both functions as goroutines
    go incrementing(c, done)
    
    go decrementing(c, done)


    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    //time.Sleep(500*time.Millisecond)
    

   
    <-done
    <-done

    i :=<-c

    fmt.Println("The magic number is:", i)
}
