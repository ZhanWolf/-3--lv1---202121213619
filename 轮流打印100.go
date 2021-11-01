package main

import (
	"fmt"
	"sync"
)

var ch=make(chan int,1)
var wg1 sync.WaitGroup


func add2()  {

	for i:=0;i<100;i++ {
		<-ch
		fmt.Println(i+1)
		ch<-0
	}
	wg1.Done()

}

func main(){
	wg1.Add(2)
	ch<-0
	go add2()
    go add2()
    wg1.Wait()
}