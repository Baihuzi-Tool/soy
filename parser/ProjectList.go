package parser

import (
	"regexp"
	"soy/config"
	"soy/engine"
	"soy/model"
)

const projectRe = `<a href="([/a-zA-Z]+)"[^>]*><span>([^>]*)<`

func ProjectListParserFun(content []byte) engine.ParserResult {

	re := regexp.MustCompile(projectRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	for _, sub := range submatch {
		url := config.BaseUrl + string(sub[1])
		request := engine.Request{
			Url:        url,
			ParserFunc: NilParserFun,
		}
		result.Requests = append(result.Requests, request)

		item := model.Project{
			Name: string(sub[2]),
			Url:  url,
		}

		result.Items = append(result.Items, item)
	}

	return result
}
