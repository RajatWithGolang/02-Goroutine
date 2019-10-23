package main

import (
	"fmt"
	"time"
)
func getChar(s string){
     for _,value := range s{
		 fmt.Printf("%c\n",value)
		 time.Sleep(1 * time.Millisecond)
	 }
}
func getDigits(a []int){
   for _,value := range a{
		 fmt.Printf("%d\n",value)
		 time.Sleep(1 * time.Millisecond)
	 }
	 fmt.Println()
}
func main(){
	fmt.Println("Main Started")
	go getChar("Hello")
	go getDigits([]int{1,2,3,4,5})
	fmt.Println("Main is Sleeping")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main has awaken Terminated")
}

