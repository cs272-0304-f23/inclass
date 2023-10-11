package main

type DownloadResult {
	body []byte
}

type ExtractResult {
	words, hrefs []string
}

func Download(inC chan string, outC chan DownloadResult) {
	for url := range inC {
		// http.Get()...

		outC <- DownloadResult{nil}
	}
}

func Extract(inC chan DownloadResult, outC chan ExtractResult) {
	for dlres := range inC {
		// extract code here
		outC <- ExtractResult{nil, nil}
	}
}

func Crawl() {
	dInCh := make(chan string, 10)
	dOutCh := make(chan DownloadResult, 10)
	eOutCh := make(chan ExtractResult, 10)

	go Download(dInCh, dOutCh)
	go Extract(dOutCh, eOutCh)
}
