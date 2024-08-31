package main

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed baseparam.json
var baseParamFS embed.FS

// base param name -> slot -> value
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

var EquipSlotBaseParamSlot = map[int]string{
	// source: https://github.com/ackwell/gear-planner/blob/master/src/data/stat.ts#L37-L62

	// 1:  "UnderArmor",
	// 1:  "Waist",
	// 1: "ChestHeadLegsFeet",
	1:  "OneHandWeapon",
	2:  "OffHand",
	3:  "Head",
	4:  "Chest",
	5:  "Hands",
	7:  "Legs",
	8:  "Feet",
	9:  "Earring",
	10: "Necklace",
	11: "Bracelet",
	12: "Ring",
	13: "TwoHandWeapon",
	15: "ChestHead",
	18: "LegsFeet",
	19: "HeadChestHandsLegsFeet",
	20: "ChestLegsGloves",
	21: "ChestLegsFeet",
}

func BaseParamModifiers(item GearItem) Stats {
	slotName := EquipSlotBaseParamSlot[item.EquipSlot]

	return Stats{
		MainStats: MainStats{
			STR: BaseParam["Strength"][slotName],
			DEX: BaseParam["Dexterity"][slotName],
			VIT: BaseParam["Vitality"][slotName],
			INT: BaseParam["Intelligence"][slotName],
			MND: BaseParam["Mind"][slotName],
		},
		SecondaryStats: SecondaryStats{
			CRIT: BaseParam["Critical Hit"][slotName],
			DET:  BaseParam["Determination"][slotName],
			DH:   BaseParam["Direct Hit Rate"][slotName],
			SKS:  BaseParam["Skill Speed"][slotName],
			SPS:  BaseParam["Spell Speed"][slotName],
			TNC:  BaseParam["Tenacity"][slotName],
			PT:   BaseParam["Piety"][slotName],
		},
	}
}
