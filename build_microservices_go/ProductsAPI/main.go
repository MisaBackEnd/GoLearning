package main

import (
	"ProductsAPI/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// create handlers
	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	// create server mux
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create server
	s := http.Server{
		Addr:         ":9090",           // the bind address
		Handler:      sm,                // set the default handler, in this case our ServeMux
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// anonymus function to start the server in non blocking fashion
	go func() {
		l.Println("API Products Server started in port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server %s\n", err)
			os.Exit(1)
		}
	}()

	// gracefull termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Received a terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
