package main

import (
	"embed"
	"encoding/json"
	"log"
)

// JSONs produced by scraping Eorzea Database (with eorzea_spider.py) miss some items (for example on August 11, 2024 Resilient gear was still hidden and marked with ??? (probably to avoid spoilers?))

//go:embed items-xivapi.json
var f embed.FS

type GearDB struct {
	// job -> category -> name
	Index map[string]map[string][]string
	Items map[string]GearItem
}

var Gear GearDB

func init() {
	Gear = LoadGearJSON()
}

func LoadGearJSON() GearDB {
	data, err := f.ReadFile("items-xivapi.json")
	if err != nil {
		log.Fatal(err)
	}

	var gear GearDB
	err = json.Unmarshal(data, &gear)
	if err != nil {
		log.Fatal(err)
	}

	return gear
}

func (db GearDB) Item(name string) GearItem {
	v, ok := db.Items[name]
	if ok {
		v.Name = name
		return v
	}

	panic("not found " + name)
}

// // job -> item type -> item name -> item stats
// type GearDB map[string]map[string]map[string]GearItem

// func LoadGearJSON() GearDB {
// 	data, err := f.ReadFile("items.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var gear GearDB
// 	err = json.Unmarshal(data, &gear)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return gear
// }

// func (db GearDB) Item(name string) GearItem {
// 	for job := range db {
// 		for category := range db[job] {
// 			v, ok := db[job][category][name]
// 			if ok {
// 				fmt.Printf("found %s in %s: %s\n", name, job, category)
// 				return v
// 			}
// 		}
// 	}

// 	panic("not found " + name)
// }
