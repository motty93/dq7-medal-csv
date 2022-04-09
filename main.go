package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var (
	records = [][]string{}
)

func main() {
	doc, err := goquery.NewDocument("https://kyokugen.info/dq7/dq7_medal.html")
	if err != nil {
		log.Fatalln(err)
		return
	}

	doc.Find("table.table1").Children().Find("tr").Each(func(i int, s1 *goquery.Selection) {
		s1.Find("td").Each(func(j int, s2 *goquery.Selection) {
			fmt.Println(j, s2.Text())
		})
	})
}
