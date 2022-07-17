package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

const (
	publicURL = "https://www.amazon.co.uk/hz/wishlist/printview/2NCDBL3HU5EGK?target=_blank&ref_=lv_pv&filter=unpurchased&sort=default"
)

type Wishlist struct {
	Items []Item
}

type Item struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	wishlist := Wishlist{}

	c := colly.NewCollector()

	c.OnHTML("tr.g-print-view-row", func(e *colly.HTMLElement) {
		parts := strings.Split(e.ChildText(".a-align-center"), "by ")
		authorParts := strings.Split(parts[1], " (")

		wishlist.Items = append(wishlist.Items, Item{
			Title:  e.ChildText(".a-text-bold"),
			Author: authorParts[0],
		})
	})

	c.Visit(publicURL)

	file, _ := json.MarshalIndent(wishlist.Items, "", " ")

	_ = ioutil.WriteFile("books.json", file, 0644)
}
