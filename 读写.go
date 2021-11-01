package main

import (
	"fmt"
	"os"
)

func main(){

	file,err := os.Create("plan.txt")
	if err !=nil {
		fmt.Println("err",err)
	}
	p:=[]byte("I’m not afraid of difficulties and insist on learning programming")
	_,err =file.Write(p)

	if err !=nil {
		fmt.Println("err",err)
	}else {
		fmt.Println("写入成功")
	}
	temp:=make([]byte,128)
	count,_:=file.ReadAt(temp,0)
    fmt.Println(string(temp[:count]))

}
