package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool){
	words:=[]string{"The","Quick","brown","fox"}
	i:=0
	defer close(wordChannel)

	t:=time.NewTimer(3* time.Second)

	for{
		select {
		case wordChannel<-words[i]  :
			i+=1
			if i==4{
				i=0
			}
        case <-done:
        	done <-true
			return
        case <-t.C:
			return
		}
	}
}



func main() {
	wordChan:=make(chan string)
	doneCh:=make(chan bool)

	go emit(wordChan,doneCh)

	for word :=range wordChan{
      fmt.Printf("%s\n",word)
	}


}
