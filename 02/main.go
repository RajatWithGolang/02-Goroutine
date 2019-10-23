package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Main is sleeping")
	fmt.Println("Hello")
}

func main(){
	fmt.Println("Main Started")
	fmt.Println("Calling Hello")
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main has awoken and Terminated")
	

}