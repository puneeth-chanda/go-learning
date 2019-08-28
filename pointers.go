package main

import "fmt"

func main(){
	a := 12
	x := &a //memory adress
	fmt.Println(x)
	*x = 5
	fmt.Println(*x)
}