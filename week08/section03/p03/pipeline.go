package main

type DownloadResult struct {
	body []byte
}

type ExtractResult struct {
	words, hrefs []string
}

func Download(inC chan string, outC chan DownloadResult) {
	// Range blocks when there's nothing to do.
	// Range stops when the channel is closed
	// Use range to monitor one channel (vs multiple with select)
	for url := range inC {
		// http.Get() ...

		outC <- DownloadResult{nil}
	}
}

func Extract(inC chan DownloadResult, outC chan ExtractResult) {
	for dl := range inC {
		// extract code goes here
		outC <- ExtractResult{nil, nil}
	}
}

func Crawl(url string) {
	dlInC := make(chan string, 10)
	dlOutC := make(chan DownloadResult, 10)
	exOutC := make(chan ExtractResult, 10)

	dlInC <- url

	go Download(dlInC, dlOutC)
	go Extract(dlOutC, exOutC)
}
