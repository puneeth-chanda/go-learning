package main

import(
	"fmt"
	"time"
)
var start time.Time

func init(){
	start = time.Now()
}

func getChars(s string){
	for _,c := range s{
		fmt.Printf("%c char at time",c, time.Since(start))
		time.Sleep(10*time.Millisecond)
	}
}

func getDigits(a []int){
	for _,d := range a{
		fmt.Printf("%d at time",d,time.Since(start))
		time.Sleep(10*time.Millisecond)
	}
}
func main(){
	fmt.Println("main start", time.Since(start))
	go getChars("Hello")
	go getDigits([]int{1,2,3,4})
	time.Sleep(100*time.Millisecond)
}