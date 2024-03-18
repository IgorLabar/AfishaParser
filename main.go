package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Concert struct {
	Title string
	Date  []string
	Place string
	Genre string
	Link  string
}

func main() {
	pages := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	concerts := Concert{}
	const fileName string = "concerts.csv"

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Не удалось создать файл: %s %s", fileName, err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Название", "Дата", "Место", "Жанр", "Ссылка"})

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	})

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

			concerts.Title = title
			concerts.Date = date[:len(date)-1]
			concerts.Place = place
			concerts.Genre = genre
			concerts.Link = link

			writer.Write([]string{concerts.Title, strings.Join(concerts.Date, ""), concerts.Place, concerts.Genre, "https://www.afisha.ru" + concerts.Link})
		}
	})

	for _, page := range pages {
		if page == 0 {
			c.Visit("https://www.afisha.ru/msk/standup/")
		} else {
			c.Visit(scrapeUrl(page))
		}
	}
}

func scrapeUrl(x int) string {
	return fmt.Sprintf("https://www.afisha.ru/msk/standup/page%v/", x)
}
