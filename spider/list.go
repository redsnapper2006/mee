package spider

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"rs.mee/crawl/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// FetchList is to fetch all company data id list
func FetchList() []string {

	request := gorequest.New()
	resp, _, _ := request.Get(utils.MeeHost + utils.MeeRootURL).End()
	// fmt.Println(body, errs)
	// fmt.Println(errs)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	s := doc.Find("body > div.content > div:nth-child(5) > div")
	re := regexp.MustCompile(`共(\d+)页`)
	pageCountArr := re.FindStringSubmatch(strings.TrimSpace(s.Nodes[0].FirstChild.Data))
	pageCount, _ := strconv.ParseInt(pageCountArr[1], 10, 32)
	fmt.Println(pageCount)

	var links []string
	table := doc.Find("body > div.content > div.content-result.clear_float > div > table > tbody")
	table.Children().Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}
		t := row.Children().Last().Children().Last().AttrOr("href", "no url")
		links = append(links, t)
		fmt.Println(i, t)
	})

	for i := 2; i <= int(pageCount); i++ {
		request := gorequest.New()
		resp, _, _ := request.Post(utils.MeeHost+utils.MeeRootURL).
			Set("Content-Type", "application/x-www-form-urlencoded").
			Send(`{"page.pageNo":"` + strconv.FormatInt(int64(i), 10) + `"}`).
			// SetCurlCommand(true).
			End()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		table := doc.Find("body > div.content > div.content-result.clear_float > div > table > tbody")
		table.Children().Each(func(i int, row *goquery.Selection) {
			if i == 0 {
				return
			}
			t := row.Children().Last().Children().Last().AttrOr("href", "no url")
			links = append(links, t)
			fmt.Println(i, t)
		})
	}
	return links
}
