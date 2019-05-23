package main

import (
	"go-spider/engine"
	"go-spider/parser"
)

func main() {
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.PrintCityList,
	})
}
