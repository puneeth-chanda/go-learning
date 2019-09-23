package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibo() func() int {
	a,b,fake := 0,1,0
	return func() int {
		fake =a
		b = a+b
		a=b
		return fake
	}
}
func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
