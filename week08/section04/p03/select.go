package main

import (
	"fmt"
)

type DownloadResult struct {
	body []byte
}

func Download(url string, ch chan DownloadResult) {
	// http.get, io.ReadAll

	ch <- DownloadResult{nil}
}

func Crawl(url string) {
	// make a channel to write URLs into for downloading
	// here, 10 specifies a buffered channel which can hold
	// 10 items before it blocks
	dlInC := make(chan string, 10)
	dlOutC := make(chan DownloadResult, 10)

	// put the seed URL into the channel
	dlInC <- url

outer:
	for {
		// use select to monitor a group of channels for readiness
		// in this example, fire off short-lived goroutines for each task
		select {
		case url := <-dlInC:
			go Download(url, dlOutC)
		case dlRes := <-dlOutC:
			go Extract(dlRes)
		case quit := <-quitC:
			// break out of the for look as well as the select
			break outer
		}
	}
	close(dlInC)
	close(dlOutC)
}
