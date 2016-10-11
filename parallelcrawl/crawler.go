package main

import (
	"fmt"
	"sync"

	"github.com/ygpravikumar/tutorial/safecounter"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var counter = safecounter.CreateCounter()

func Crawl(url string, depth int, fetcher Fetcher) {
	var str_map = make(map[string]bool)
	var mux sync.Mutex
	var wg sync.WaitGroup

	var crawler func(string, int)
	crawler = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 {
			return
		}

		mux.Lock()
		if _, ok := str_map[url]; ok {
			mux.Unlock()
			return
		} else {
			str_map[url] = true
			mux.Unlock()
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q %q\n", url, body, urls)

		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1)
		}
	}
	wg.Add(1)
	crawler(url, depth)
	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
