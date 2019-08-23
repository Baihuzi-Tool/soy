package parser

import "soy/engine"

func NilParserFun(content []byte) engine.ParserResult {
	return engine.ParserResult{}
}
