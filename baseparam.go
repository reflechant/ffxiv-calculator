package main

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed baseparam.json
var baseParamFS embed.FS

var BaseParam map[string]map[string]int

func LoadBaseParam() map[string]map[string]int {
	data, err := baseParamFS.ReadFile("baseparam.json")
	if err != nil {
		log.Fatal(err)
	}

	var bp map[string]map[string]int
	err = json.Unmarshal(data, &bp)
	if err != nil {
		log.Fatal(err)
	}

	return bp
}
