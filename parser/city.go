package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kentio/cmd-health-examination/config"
	"github.com/kentio/cmd-health-examination/engine"
	"github.com/kentio/cmd-health-examination/model"
	"log"
	"regexp"
	"strings"
)

const cityRe = `<a href="(care-[0-9a-z]+.html)">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	var city string

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil{
		log.Fatalln(err)
	}
	// curtury city
	dom.Find("strong").Each(func(i int, selection *goquery.Selection) {
		city = selection.Text()
		fmt.Println(i, city)
	})
	dom.Find("li").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("h3").Text()
		//content := strings.TrimSpace(selection.Find("p").Text())
		content := ""
		selection.Find("p").Each(func(i int, selection *goquery.Selection) {
			content += selection.Text() + "\n"
		})

		config.DataDB.Create(&model.Hospital{City: city, Name: title, Content: content})
	})


	// xxx
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: config.BaseUrl + string(m[1]), ParserFunc: engine.NilParser})
		fmt.Printf("%s %s\n", m[1], m[2])
	}
	fmt.Println(len(matches))
	return result
}
