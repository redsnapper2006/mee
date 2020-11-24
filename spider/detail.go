package spider

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

func FetchDetail(url string) []string {

	request := gorequest.New()
	resp, _, _ := request.Get(url).End()
	// fmt.Println(body, errs)
	// fmt.Println(errs)
	// fmt.Println(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(doc.Text())
	table := doc.Find("body > div:nth-child(3) > div:nth-child(1) > table.tab0 > tbody")
	// fmt.Println(table.Text())
	table.Children().Each(func(i int, row *goquery.Selection) {
		// fmt.Println(row)
		if i == 0 {
			return
		}
		row.Children().Each(func(j int, subRow *goquery.Selection) {
			fmt.Println(subRow.Nodes[0].FirstChild.Data)
		})
	})

	table = doc.Find("body > div:nth-child(3) > div:nth-child(1) > table:nth-child(4) > tbody")
	// fmt.Println(table.Text())
	table.Children().Each(func(i int, row *goquery.Selection) {
		// fmt.Println(row)
		row.Children().Each(func(j int, subRow *goquery.Selection) {
			fmt.Println(subRow.Children().Last().AttrOr("href", "no url"))
		})
	})


//	body > div:nth-child(2) > table > tbody
	return nil
}
