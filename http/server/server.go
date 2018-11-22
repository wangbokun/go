package main

import (
	"net/http"
)


func hello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello word!"))
}


func main(){
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":8001",nil)
}