package main

import (
	"as/internal/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	/* build routing */
	router := http.NewServeMux()
	for _, route := range routes.Routes {
		router.HandleFunc(route.Url, routes.Handler(route.Func, route.Method))
	}

	/* create server */
	server := &http.Server{
		Addr:         ":9001",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	/* start server */
	log.Println("Server started at http://localhost:9001")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
