package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
	"unicode"
)

type Link struct {
	Href string
	Text string
}

func FindLinks(f io.Reader, domain string) []Link {
	z := html.NewTokenizer(f)
	depth := 0
	var link Link
	var links []Link

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
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
						if strings.Contains(link.Href, domain) {
							links = append(links, link)
						}
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
