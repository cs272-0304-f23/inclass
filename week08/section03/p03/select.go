package main

type DownloadResult struct {
	body []byte
	err  error
	url  string // may be needed by subsequent goroutines
}

func Download(wg *sync.WaitGroup, url string, outC chan DownloadResult) {
	wg.Add(1)
	defer wg.Done()

	// http.Get()...
	outC <- DownloadResult{nil, nil}
}

func Extract(dl *DownloadResult) {

}

func Crawl(url string) {

	// A channel without a length is synchronous
	// Data can be written when the channel is empty
	// if channel has data in it, writes to the channel block

	// If we give the channel a length, that specifies a buffered
	// channel, where we can write 10 items in without blocking
	// Writing the 11th item blocks
	dlInC := make(chan string, 10)
	dlOutC := make(chan DownloadResult, 10)

	var wg sync.WaitGroup

outer:
	for {
		// Select is like switch, but for channels. It cycles around
		// a set of channels, running the code for a ready channel
		select {
		case url := <-dlInC:
			go Download(url, dlOutC)
		case dl := <-dlOutC:
			go Extract(&dl)
		case q := <-quit:
			// If we knew how to tell when we're done, a quit channel
			// is a common Go idiom for breaking inner/outer loops
			// using break <label>
			break outer
		}

		// I tried this but channels can be zero while there is
		// still work outstanding
		if len(dlInC) == 0 && len(dlOutC) == 0 {
			break
		}

		go func() {
			wg.Wait()
			quit <- true
		}()
	}

	close(dlInC)
	close(dlOutC)
}
