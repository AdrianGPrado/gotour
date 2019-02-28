package main

import (
	"fmt"
	"sync"
)

// Fetcher something
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Fetched store fetched urls to prevent double fetching.
type Fetched struct {
	counter int
	urls    map[string]bool
	mux     sync.Mutex
}

func (f *Fetched) hasURL(url string) bool {
	if has, ok := f.urls[url]; ok {
		return has
	}
	return false
}

func (f *Fetched) addURL(url string) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.urls[url] = true
}

func (f *Fetched) addCounter() {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.counter++
}

func (f *Fetched) subtractCounter() {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.counter--
}
func (f *Fetched) printURLs() {
	fmt.Printf("urls proccessed:\n")
	for i, u := range f.urls {
		fmt.Printf("%v %v\n", i, u)
	}
}

var fetched = Fetched{urls: make(map[string]bool), counter: 1}

// Link keep together a url and it's depth in the graph
type Link struct {
	url   string
	depth int
}

func crawler(l Link, fetcher Fetcher, links chan Link) {
	fetched.subtractCounter()
	fmt.Printf("processing %v, counter %v\n", l, fetched.counter)
	_, urls, err := fetcher.Fetch(l.url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fetched.addURL(l.url)
	if fetched.counter == 0 && l.depth == 2 {
		fmt.Printf("exit conditions : %v depth: %v \n", fetched.counter, l.depth)
		close(links)
	} else {
		for _, u := range urls {
			links <- Link{u, l.depth + 1}
		}
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	links := make(chan Link)
	// done := make(chan string)

	go crawler(Link{url: url, depth: 0}, fetcher, links)

	for link := range links {
		fmt.Printf("received %v, counter %v\n", link, fetched.counter)
		if !fetched.hasURL(link.url) && link.depth < depth {
			fetched.addCounter()
			go crawler(link, fetcher, links)
		}
	}
	// fetched.printURLs()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
