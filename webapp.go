package main

import("fmt";
		"net/http")
func Indexhandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Go is great") //prints by formatting
}
//similar functions can be written
func main(){

	http.HandleFunc("/", Indexhandler)
//http.HandleFunc("/about/",AboutHandler)
	http.ListenAndServe(":8000", nil)
}
