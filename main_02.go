package main

import (
	"os"
	"runtime/pprof"
)

func main() {

	f,err:=os.Create("./cpu.prof")
	if err!=nil{

		panic(err)
	}
	if err:=pprof.StartCPUProfile(f);err!=nil{

		panic(err)
	}
	defer pprof.StopCPUProfile()
}
