package parser

import (
	"fmt"
	"github.com/kentio/cmd-health-examination/config"
	"github.com/kentio/cmd-health-examination/engine"
	"regexp"
)

const cityListRe = `<a href="(care-[0-9a-z]+.html)">([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: config.BaseUrl + string(m[1]),
				ParserFunc: ParseCity})
		fmt.Printf("%s %s\n", m[1], m[2])
	}
	fmt.Println(len(matches))
	return result
}
