package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkPageStatus(pageUrl string, ch chan string) {
	_, err := http.Get(pageUrl)

	if err != nil {
		ch <- pageUrl + " is Down"
	} else {
		ch <- pageUrl + " is Up"
	}
}

func main() {
	start := time.Now()
	channel := make(chan string)
	servidores := []string{
		"https://mesfix.com",
		"https://google.com",
		"https://plataform.com",
		"https://misaelriot.tv",
	}

	for _, servidor := range servidores {
		go checkPageStatus(servidor, channel)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-channel)
	}
	enlapsedTime := time.Since(start)
	fmt.Printf("Enlapsed Time:", enlapsedTime)
}
