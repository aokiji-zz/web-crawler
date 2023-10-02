package links

import (
	"fmt"
	"net/url"
	"time"
	"web-crawler/db"

	"golang.org/x/net/html"
)

type VisitedLinkType struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visitedDate"`
}

var links []string

func ExtractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}

			if db.VisitedLink(link.String()) {
				fmt.Printf("link visitado: %s\n", link)
				continue
			}
			visitedLink := VisitedLinkType{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}
			db.Insert("links", visitedLink)
			links = append(links, link.String())
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractLinks(c)
	}
}
