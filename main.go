package main

import (
	"flag"
	"log"
	"fmt"
	
	"github.com/parnurzeal/gorequest"
)

var (
	userId = flag.String("UserId", "imagineer", "Set target brunch user id")
	urlMagazine string
	urlArticle string
)

func init() {
	flag.Parse()
	
	urlMagazine = fmt.Sprintf("https://api.brunch.co.kr/v2/magazine/@%s", *userId)
	urlArticle	= fmt.Sprintf("https://api.brunch.co.kr/v1/article/@%s", *userId)
}

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get(urlMagazine).End()

	if len(errs) == 0 {
		log.Printf("resp: %s", resp)
		log.Printf("body: %s", body)
	}
}
