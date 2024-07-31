package main

import (
	"github.com/gocolly/colly"
)

// store in struct
type item struct {
	Name string `json:"name"`
}

func main() {
	c := colly.NewCollector()
	//store in list
	var items []item
	c.OnHTML("div[data-cy=title-recipe]", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("h2>a>span"),
		}

		items = append(items, item)
	})

	c.Visit("https://www.amazon.in/s?k=best+smartphone+under+30000&crid=15PX311LNCWD&sprefix=best+smartphone+%2Caps%2C264&ref=nb_sb_ss_pltr-xclick_6_16")
}
