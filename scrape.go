package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Listing struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Location string `json:"location"`
	Link string `json:"link"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// set an empty slice of listing structs
	listings := []Listing{}

	c.OnHTML("li.cl-static-search-result", func(e *colly.HTMLElement) {
		// scrape the data
		title := e.ChildText("div.title")
		price := e.ChildText("div.price")
		location := e.ChildText("div.location")
		link := e.ChildAttr("a", "href")
		// create a new listing struct
		listing := Listing{title, price, location, link}
		// append the listing struct to the slice
		listings = append(listings, listing)
	})

	c.OnScraped(func(r *colly.Response) {
		// Once everything is scraped, turn the slice into a JSON string
		listingsJSON, err := json.Marshal(listings)
		if err != nil {
			log.Fatalf("Failed to encode JSON: %v", err)
	}
	// create the file
	file, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// add the JSON to the file
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(listingsJSON))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://vancouver.craigslist.org/search/apa")
}
