package parser

import (
	"ericivan/crawler/engine"
	"regexp"
)

const citiListRegex = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)" [^>]+>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {

	result := engine.ParseResult{}
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)" [^>]+>([^<]+)</a>`)

	res := re.FindAllSubmatch(content, -1)

	limit := 1
	for _, val := range res {
		result.Request = append(result.Request, engine.Request{
			Url:        string(val[1]),
			ParserFunc: ParseCity,
		})

		result.Items = append(result.Items, string(val[2]))

		limit--

		if limit == 0 {
			break
		}
	}

	return result
}
