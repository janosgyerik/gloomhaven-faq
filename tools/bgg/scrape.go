package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly"
)

type article struct {
	XMLName xml.Name `xml:"article"`
	Subject string   `xml:"subject"`
	Body    string   `xml:"body"`
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
	})

	var url string
	// url = "https://boardgamegeek.com/thread/1897763/official-faq-game-no-rules-questions-please"
	// url = "https://api.geekdo.com/xmlapi2/thread?id=1897763"
	// requires file transport http.Dir("."), and the file already fetched to current dir
	// TODO fetch from url param
	url = "file:///thread.xml"
	c.Visit(url)
}
