
package main

import "fmt"
import "time"


func producer(buffer chan int, done chan bool){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        buffer <- i
    }
    done <- true
}

func consumer(buffer chan int, done chan bool){

    time.Sleep(1 * time.Second)
    for {
        i := <- buffer //TODO: get real value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
        //buffer <- i
    }
    
}


func main(){
    
    // TODO: make a bounded buffer
    buffer := make(chan int,1)
    done := make(chan bool)
    
    go consumer(buffer, done)
    go producer(buffer, done)
    
    //select {}
    <-done
}