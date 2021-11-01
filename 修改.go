package main

import "fmt"

func main() {

	over := make(chan bool)



		go func() {
			for i := 0; i < 10; i++ {//for循环属于主协程中，而主协程程中存在一个待接收的channel

			fmt.Println("我是协程！！",i)



		if i == 9 {

			over <- true

		}

	    }
		}()
	<-over

	fmt.Println("over!!!")

}
