package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	records = [][]string{
		{"No", "Place", "Time", "Detail"},
	}
)

func main() {
	doc, err := goquery.NewDocument("https://kyokugen.info/dq7/dq7_medal.html")
	if err != nil {
		log.Fatalln(err)
		return
	}

	doc.Find("table.table1").Children().Find("tr").Each(func(i int, s1 *goquery.Selection) {
		record := make([]string, 4)
		s1.Find("td").Each(func(j int, s2 *goquery.Selection) {
			record[j] = strings.TrimSpace(s2.Text())
		})
		records = append(records, record)
	})

	// 出力するfile作成
	file, err := os.Create("dq7_medal.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	cw := csv.NewWriter(file)
	defer cw.Flush()
	// 標準出力する場合(確認用)
	// w := csv.NewWriter(os.Stdout)

	// 予めrecordsが2次元配列の形であればWriteAllでおｋ
	cw.WriteAll(records)
	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatalln("error writing record to csv:", err)
	// 	}
	// }

	if err := cw.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	fmt.Println("csv read ok")
}
