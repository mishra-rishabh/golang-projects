package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type quote struct {
	Text   string   `json:"text"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func main() {
	/* Collector */
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	var allQuotes []quote

	// c.Visit("quotes.toscrape.com")
	// fmt.Println("hello world")

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		var tags []string

		tempQuote := quote{
			Text:   h.ChildText("span.text"),
			Author: h.ChildText("small.author"),
			Tags:   append(tags, h.ChildText("div.tags")),
		}
		// fmt.Println(tempQuote)

		allQuotes = append(allQuotes, tempQuote)
	})

	c.OnHTML("li.next > a", func(h *colly.HTMLElement) {
		nextPage := h.Attr("href")

		if nextPage != "" {
			c.Visit("https://quotes.toscrape.com" + nextPage)
		}

		fmt.Println("https://quotes.toscrape.com" + nextPage)
	})

	err := c.Visit("https://quotes.toscrape.com")

	if err != nil {
		log.Fatal(err)
	}

	content, err := json.Marshal(allQuotes)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("famous_quotes.json", content, 0644)
}
