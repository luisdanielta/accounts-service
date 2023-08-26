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
		Addr:         ":9015",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	/* start server */
	log.Println("Server started at http://localhost:9015")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
