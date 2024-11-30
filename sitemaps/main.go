package main

// > Data Types
// - One set to maintain all the links
// - One original domain
// - Temporary set to fetch all the links from the current page
// - One queue to store all the links which are not explored

// > First find the way we can find html from the given url
// > Find all url from the html
// > Filter out all the urls which are not belongs to the domain
// > Add the links in the queue if an only if it is not in the set

import (
	"fmt"
	"github.com/rajatnai49/sitemaps/link"
	"log"
	"net/http"
)

type queue []string

func (q *queue) Push(v string) {
	*q = append(*q, v)
}

func (q *queue) Pop() string {
    if q.IsEmpty() {
        return ""
    }
    first := (*q)[0]
    *q = (*q)[1:]
    return first
}

func (q *queue) IsEmpty() bool {
    return len(*q) == 0
}

type Link struct {
	Href string
	Text string
}

func main() {
	linksSet := map[string]struct{}{}
	q := &queue{}
	startUrl := "https://gophercises.com/"
	domain := "https://gophercises.com/"

    q.Push(startUrl)

    for !q.IsEmpty() {
        currentUrl := q.Pop()
        if _, visited := linksSet[currentUrl]; visited {
            continue
        }
        linksSet[currentUrl] = struct{}{}
        processPage(currentUrl, q, linksSet, domain)
    }
}

func processPage(hrefLink string, q *queue, lSets map[string]struct{}, domain string) {
    fmt.Println("Processing: ", hrefLink)
	res, err := http.Get(hrefLink)
	if err != nil {
		log.Fatal(err)
        return
	}
    defer res.Body.Close()

	ans := link.FindLinks(res.Body, domain)
    fmt.Println(ans)

    for _, item := range ans {
		href := item.Href
		if _, ok := lSets[href]; !ok {
            q.Push(href)
		}
	}
}
