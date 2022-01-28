package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/_wedding/submit", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			rw.WriteHeader(200)
			return
		}
		rw.WriteHeader(200)
		if err := r.ParseForm(); err != nil {
			rw.WriteHeader(500)
			return
		}
		fmt.Printf("received form %+v \n", r.PostForm)
	})
	fmt.Println("listening")
	if err := http.ListenAndServe("0.0.0.0:1906", nil); err != nil {
		panic(err)
	}
}