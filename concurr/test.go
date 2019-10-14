package main

import (
	"time"
	"fmt"
)
func printMe(){
	time.Sleep(15*time.Millisecond)
	fmt.Println("hello again world!!")
}

func main(){
	fmt.Println("main() started")
	go printMe()
	time.Sleep(10*time.Millisecond)
	fmt.Println("main() stopped")

}

