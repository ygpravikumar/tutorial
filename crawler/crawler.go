package main

import (
	"fmt"
	"time"

	"github.com/ygpravikumar/tutorial/safecounter"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var counter = safecounter.CreateCounter()

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ret chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer close(ret)

	if depth <= 0 {
		return
	}

	if counter.Value(url) != 0 {
		return
	}

	counter.Inc(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ret <- fmt.Sprintf(err.Error())
		return
	}

	ret <- fmt.Sprintf("found: %s %q\n", url, body)

	results := make([]chan string, len(urls))
	for i, u := range urls {
		go Crawl(u, depth-1, fetcher, results[i])
	}

	for i := range results {
		for s := range results[i] {
			ret <- s
		}
	}

	//Do not know why wihout this line nothing works
	time.Sleep(10 * time.Millisecond)
	return
}

func main() {

	ret := make(chan string)
	go Crawl("http://golang.org/", 4, fetcher, ret)
	for i := range ret {
		fmt.Printf(i)
	}

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
