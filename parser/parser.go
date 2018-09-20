package parser

import (
	"encoding/json"
	"strings"
)

func ParseTextIntoJson(text []byte) map[string]interface{} {
	lines := strings.Split(string(text), "\n")
	if strings.HasPrefix(lines[0], "//") {
		lines = lines[1:]
	}
	target := []byte(strings.Join(lines, "\n"))
	var data map[string]interface{}
	if err := json.Unmarshal(target, &data); err != nil {
        panic(err)
    }
	return data
}
