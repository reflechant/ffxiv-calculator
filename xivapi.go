package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// XIVAPIGearItem is used to unmarshal XIVAPI JSON responses
type XIVAPIGearItem struct {
	Name          string  `json:"name"`
	Lvl           uint    `json:"ilvl"`
	Jobs          Job     // bitmask
	JobLvl        uint    `json:"job level"`
	STR           int     `json:"Strength,omitempty"`
	DEX           int     `json:"Dexterity,omitempty"`
	VIT           int     `json:"Vitality,omitempty"`
	INT           int     `json:"Intelligence,omitempty"`
	MND           int     `json:"Mind,omitempty"`
	CRIT          int     `json:"Critical Hit,omitempty"`
	DET           int     `json:"Determination,omitempty"`
	DH            int     `json:"Direct Hit Rate,omitempty"`
	SKS           int     `json:"Skill Speed,omitempty"`
	SPS           int     `json:"Spell Speed,omitempty"`
	TNC           int     `json:"Tenacity,omitempty"`
	PT            int     `json:"Piety,omitempty"`
	PhysDMG       float64 `json:"Physical Damage,omitempty"`
	MagDMG        float64 `json:"Magic Damage,omitempty"`
	AutoAtk       float64 `json:"Auto-attack,omitempty"`
	Delay         float64 `json:"Delay,omitempty"`
	MateriaSlots  int     `json:"materia slots,omitempty"`
	MateriaMelded []Materia
}

var GearCategories = []string{
	"Shield", "Head", "Body", "Legs", "Hands", "Feet", "Necklace", "Earrings", "Bracelets", "Ring",
}

func LoadGear(minLvl, maxLvl int, job Job, cat string) []GearItem {
	limit := 1000
	query := fmt.Sprintf("LevelItem>=%d LevelItem<=%d ClassJobCategory.%s=true ItemUICategory.Name=%s", minLvl, maxLvl, jobNames[job], cat)
	fields := "Name,LevelEquip,LevelItem.value,DamagePhys,DamageMag,Delayms,BaseParam[].Name,BaseParamValue,ItemUICategory.Name,MateriaSlotCount"

	resp, err := http.Get(fmt.Sprintf(
		"https://beta.xivapi.com/api/1/search?sheets=Item&query=%s&limit=%d&fields=%s",
		url.QueryEscape(query),
		limit,
		url.QueryEscape(fields)))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(body)

	var gear []GearItem
	err = json.Unmarshal(body, &gear)
	if err != nil {
		log.Fatal(err)
	}

	return gear
}
