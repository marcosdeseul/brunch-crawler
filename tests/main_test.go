package main_test

import (
	"testing"
	"io/ioutil"

	"github.com/marcosdeseul/brunch-crawler/parser"
)

func TestJsonParse(t *testing.T) {
	body, err := ioutil.ReadFile("samples/article.json")
	if err != nil {
		t.Error("there is an io error: ", err)
	}

	json := parser.ParseTextIntoJson(body)

	if json["desc"] != "OK" {
		t.Error("desc of article.json should be OK", json["desc"])
	}
}
