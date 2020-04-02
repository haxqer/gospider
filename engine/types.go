package engine

import "git.trac.cn/nv/spider/model"

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

//type Item struct {
//	URL     string
//	ID      string
//	Payload interface{}
//}

type Item = model.Mgtv
