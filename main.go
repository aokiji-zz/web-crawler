package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"web-crawler/db"

	"golang.org/x/net/html"
)

var links []string

func main() {
	// url := "https://aprendagolang.com.br"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic((err))
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	panic(fmt.Sprintf("Status diferente de 200: %d", resp.StatusCode))
	// }
	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// extractLinks(doc)
	visitedLink("https://aprendagolang.com.br")

}
func visitedLink(url string) {
	// url := "https://aprendagolang.com.br"
	resp, err := http.Get(url)
	if err != nil {
		panic((err))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status diferente de 200: %d", resp.StatusCode))
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	extractLinks(doc)
}

type VisitedLink struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visitedDate"`
}

func extractLinks(node *html.Node) {
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
			visitedLink := VisitedLink{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}
			db.Insert("links", visitedLink)
			// db.FindMany("links", visitedLink)
			links = append(links, link.String())
			// fmt.Println(link.String())
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}
