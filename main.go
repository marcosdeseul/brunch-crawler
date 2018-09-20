package main

import (
	"flag"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
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

func requestGet(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func main() {
	body, err := requestGet(urlMagazine)
	if err == nil {
		log.Printf("body: %s", body)
	} else {
		log.Printf("err: %s", err)
	}
}
