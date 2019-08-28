package main

import ("fmt")

func add(x,y float64) float64{
	return(x+y)
}
func multiple(a,b string) (string,string){
	return a,b
}

func main(){
	num1,num2  := 5.6,3.5
	w1,w2 := "a","b"
	var a int = 12
	var b float64 = float64(a)
	fmt.Println(b)
	fmt.Print(add(num1,num2))
	fmt.Println(multiple(w1,w2))
}