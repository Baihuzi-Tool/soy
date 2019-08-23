package main

import (
	"soy/config"
	"soy/engine"
	"soy/parser"
)

func main() {
	request := engine.Request{
		Url:        config.BaseUrl + config.ProjectListBasePath,
		ParserFunc: parser.ProjectListParserFun,
	}
	engine.Run(request)
}
