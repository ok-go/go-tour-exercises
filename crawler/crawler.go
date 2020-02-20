package crawler

import (
	"fmt"
	"sync"
)

type fetcher interface {
	fetch(url string) (body string, urls []string, err error)
}

type mapSafe struct {
	m   map[string]byte
	mtx sync.Mutex
}

func (m *mapSafe) checkAndSet(key string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	_, ok := m.m[key]
	if !ok {
		m.m[key] = 1
	}

	return ok
}

type message struct {
	Cnt int
	Msg string
}

func (m message) Nil() bool {
	return m == message{}
}

func (m message) String() string {
	return m.Msg
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, f fetcher) {
	ch := make(chan message)
	m := &mapSafe{m: make(map[string]byte)}

	go crawl(url, depth, f, ch, m)

	var r message
	cnt := 1
	for {
		r = <-ch
		if !r.Nil() {
			fmt.Println(r)
		}

		cnt += r.Cnt - 1
		if cnt <= 0 {
			return
		}
	}
}

func crawl(url string, depth int, f fetcher, ch chan message, m *mapSafe) {
	if depth <= 0 || m.checkAndSet(url) {
		ch <- message{}
		return
	}

	body, urls, err := f.fetch(url)

	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("found: %s %q", url, body)
	}

	ch <- message{len(urls), msg}

	for _, u := range urls {
		go crawl(u, depth-1, f, ch, m)
	}

	return
}

// fakeFetcher is Fetcher that return scanned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fake = fakeFetcher{
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

func run() {
	println("11. CRAWLER:")
	Crawl("https://golang.org/", 4, fake)
}

func Run() {
	run()
}
