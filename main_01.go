package main

import (
	"fmt"
	"sync"
)

type SomeStruct struct {
	values map[string]int
	mutext sync.Mutex
}
var  wait =sync.WaitGroup{}

func(s *SomeStruct) Count(key string){

	s.mutext.Lock()
	defer s.mutext.Unlock()
	defer wait.Done()

	value,ok :=s.values[key]
	if !ok{
		s.values[key]=1
	}else{
		s.values[key]=value+1
	}

}


func main() {


	s:=&SomeStruct{
		make(map[string]int),
		sync.Mutex{},

	}
	for i:=0;i<8;i++{
		wait.Add(1)
		go s.Count("foo")
	}
	//time.Sleep(100 * time.Microsecond)
	wait.Wait()
	fmt.Printf("%d\n",s.values["foo"])
}
