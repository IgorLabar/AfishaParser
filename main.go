package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Concert struct {
	Title string
	Price string
	Genre string
	Date  string
	Place string
}

func main() {
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0"

	c.OnHTML("div.OILfh div.oP17O div.QWR1k", func(e *colly.HTMLElement) {
		selector := e.DOM
		childNodes := selector.Children().Nodes

		if len(childNodes) == 3 {
			title := selector.Find("a.CjnHd").Text()
			value := selector.FindNodes(childNodes[2]).Text()
			date := strings.Split(value, ",")
			place := date[len(date)-1]
			genre := selector.FindNodes(childNodes[1]).Text()
			link := e.ChildAttr("a.CjnHd", "href")
			price := e.DOM.Before("div.MckHJ").Text()

			fmt.Printf("Name: %s, Date: %s, Place: %s Genre: %s Link: https://www.afisha.ru%s, Price: %s \n", title, date[:len(date)-1], place, genre, link, price)
		}
	})

	c.Visit("https://www.afisha.ru/msk/standup/")
}
