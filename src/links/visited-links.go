package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func VisitedLink(url string) {
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
	ExtractLinks(doc)
}
