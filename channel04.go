package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reder(ch chan int){
	t:=time.NewTimer(3 * time.Second)
	for{
		select {
		    case i:=<-ch:
		    	fmt.Printf("%d\n",i)
		    case <-t.C:
		    	ch=nil
		}
	}
}
func writer(ch chan int){
	for{
		ch<-rand.Intn(42)
	}
}

func main() {
   ch:=make(chan int)
   go reder(ch)
   go writer(ch)
   time.Sleep(10 * time.Second)
}
