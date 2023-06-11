package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// funcion anonima
			defer func() {
				log.Println(r.Method, r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
