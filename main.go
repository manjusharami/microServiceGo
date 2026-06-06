package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
 func HandleFuncs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellow world")
}

func main() {
	router:=chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello",HandleFuncs)


	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("server is runing")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server failed to start", err)
	}
	 
}
