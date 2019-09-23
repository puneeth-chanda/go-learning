package main

import("fmt";
		"net/http")
func Indexhandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Go is great<h1>") //prints by formatting
	fmt.Fprintf(w,"<p>You can %s and %s</p>", "do","get")
	fmt.Fprintf(w,`dont care what others are feeling about you and enjoy it`)
}
//similar functions can be written
func main(){

	http.HandleFunc("/", Indexhandler)
//http.HandleFunc("/about/",AboutHandler)
	http.ListenAndServe(":8000", nil)
}
