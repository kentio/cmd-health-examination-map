package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kentio/cmd-health-examination-map/config"
	"github.com/kentio/cmd-health-examination-map/engine"
	"github.com/kentio/cmd-health-examination-map/model"
	"log"
	"strings"
)

const cityRe = `<a href="(care-[0-9a-z]+.html)">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	var city string

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatalln(err)
	}
	// data struct
	result := engine.ParseResult{}

	// curtury city
	dom.Find("strong").Each(func(i int, selection *goquery.Selection) {
		city = selection.Text()
		fmt.Println(i, city)
	})
	dom.Find("li").Each(func(i int, selection *goquery.Selection) {
		detail := model.Hospital{City: city}
		detail.Name = selection.Find("h3").Text()

		selection.Find("p").Each(func(i int, selection *goquery.Selection) {
			detail.Content += selection.Text() + "\n"
		})
		// save data to db
		fmt.Println(detail.Content)
		result.Items = append(result.Items, detail)
		config.DataDB.Create(&detail)
	})
	return result
}
