package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Provide HTML file")
        return
    }
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("There is error in the opening a file: ", err)
		return
	}
	z := html.NewTokenizer(f)
	depth := 0
	var link Link
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && 'a' == tn[0] {
				if html.StartTagToken == tt {
					if depth < 1 {
						key, ta, more := z.TagAttr()
						if "href" == string(key) {
							link.Href = string(ta)
						}
						for more {
							if "href" == string(key) {
								link.Href = string(ta)
							}
							key, ta, more = z.TagAttr()
						}
					}
					depth++
				} else {
					depth--
					if depth == 0 {
                        link.Text = strings.TrimSpace(link.Text)
						fmt.Println(link)
						link.Text = " "
						link.Href = " "
					}
				}
			}
		case html.TextToken:
			if depth > 0 {
				link.Text += strings.TrimLeftFunc(string(z.Text()), unicode.IsSpace)
			}
		}
	}
}
