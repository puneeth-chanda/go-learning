package main

import ("fmt")

func add(x,y float64) float64{
	return(x+y)
}
func main(){
	num1,num2  := 5.6,3.5
	fmt.Print(add(num1,num2))
}