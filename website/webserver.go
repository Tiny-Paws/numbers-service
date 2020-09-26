package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website/static")))
	http.ListenAndServe(":8091", nil)
}
