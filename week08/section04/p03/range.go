package main

type DownloadResult struct {
	body []byte
}

// Rather than calling a short-running goroutine from select,
// we could use a long-running goroutine using range
func Download(inC chan string, outC chan DownloadResult) {

	// range runs as long as the channel is open
	// it doesn't stop when there's nothing to do
	for url := range inC {
		// http.Get()...

		outC <- DownloadResult{nil}
	}
}

func Crawl(url string) {
	dlInC := make(chan string, 10)
	dlOutC := make(chan DownloadResult, 10)

	go Download(dlInC, dlOutC)
	dlInC <- url

}
