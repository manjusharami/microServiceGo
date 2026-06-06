package main

import (
	"fmt"
	"net/http"
	"chi"
)
 func HandleFuncs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow world")
}

func main() {

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(HandleFuncs),
	}

	fmt.Println("server is runing")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server failed to start", err)
	}
	 
}
