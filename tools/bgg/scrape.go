package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"

	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly"
)

type article struct {
	XMLName xml.Name `xml:"article"`
	Subject string   `xml:"subject"`
	Body    string   `xml:"body"`
}

func extractHeadings(html string) []string {
	before := `&lt;br/&gt;&lt;br/&gt;&lt;b&gt;`
	after := `&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;`
	r := regexp.MustCompile(fmt.Sprintf(`%s(?P<heading>.*?)%s`, before, after))
	matches := r.FindAllStringSubmatch(html, -1)
	headings := make([]string, len(matches))
	if matches != nil {
		for i, match := range matches {
			headings[i] = match[1]
		}
	}
	return headings
}

func main() {
	c := colly.NewCollector()
	// TODO fetch file to workdir, re-fetch depending on creation date
	c.WithTransport(http.NewFileTransport(http.Dir(".")))
	c.OnXML("//articles/article[1]", func(e *colly.XMLElement) {
		// TODO produce nice header
		// TODO produce content only
		// TODO produce TOC
		fmt.Println(e.Attr("id"))
		fmt.Println(e.Attr("postdate"))
		fmt.Println(e.Attr("editdate"))
		fmt.Println(e.Attr("numedits"))
		d := e.DOM.(*xmlquery.Node)
		fmt.Println(d.Data)
		fmt.Println(d.SelectElement("subject"))
		var a article
		xml.Unmarshal([]byte(d.OutputXML(true)), &a)
		fmt.Println(a.Subject)
		fmt.Println(a.Body)

		// &lt;br/&gt;&lt;br/&gt;&lt;br/&gt;&lt;b&gt;Characters&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;
		for _, heading := range extractHeadings(a.Body) {
			fmt.Println(heading)
		}
		fmt.Println("what?")
	})

	var url string
	// url = "https://boardgamegeek.com/thread/1897763/official-faq-game-no-rules-questions-please"
	// url = "https://api.geekdo.com/xmlapi2/thread?id=1897763"
	// requires file transport http.Dir("."), and the file already fetched to current dir
	// TODO fetch from url param
	url = "file:///thread.xml"
	c.Visit(url)
}
