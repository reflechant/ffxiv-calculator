package main

import (
	"math"
)

type AvailableGear struct {
	Weapon  []GearItem
	OffHand []GearItem

	Head  []GearItem
	Body  []GearItem
	Hands []GearItem
	Legs  []GearItem
	Feet  []GearItem

	Ears      []GearItem
	Neck      []GearItem
	Wrist     []GearItem
	LeftRing  []GearItem // to avoid special case of unique rings
	RightRing []GearItem
}

var SlotItemUICategory = map[string]string{
	"Weapon":    "Weapon",
	"OffHand":   "OffHand",
	"Head":      "Head",
	"Body":      "Body",
	"Hands":     "Hands",
	"Legs":      "Legs",
	"Feet":      "Feet",
	"Ears":      "Earrings",
	"Neck":      "Necklace",
	"Wrist":     "Bracelets",
	"LeftRing":  "Ring",
	"RightRing": "Ring",
}

func (ag AvailableGear) ToMap() map[string][]GearItem {
	return map[string][]GearItem{
		"Weapon":    ag.Weapon,
		"OffHand":   ag.OffHand,
		"Head":      ag.Head,
		"Body":      ag.Body,
		"Hands":     ag.Hands,
		"Legs":      ag.Legs,
		"Feet":      ag.Feet,
		"Ears":      ag.Ears,
		"Neck":      ag.Neck,
		"Wrist":     ag.Wrist,
		"LeftRing":  ag.LeftRing,
		"RightRing": ag.RightRing,
	}
}

func (ag *AvailableGear) LoadFromMap(m map[string][]GearItem) {
	ag.Weapon = m["Weapon"]
	ag.OffHand = m["OffHand"]
	ag.Head = m["Head"]
	ag.Body = m["Body"]
	ag.Hands = m["Hands"]
	ag.Legs = m["Legs"]
	ag.Feet = m["Feet"]
	ag.Ears = m["Ears"]
	ag.Neck = m["Neck"]
	ag.Wrist = m["Wrist"]
	ag.LeftRing = m["LeftRing"]
	ag.RightRing = m["RightRing"]
}

// Combinations will return all possible gear set combinations out of available gear
// For now we meld with materia 12 only
func (g AvailableGear) Combinations(materiaTypes []*Materia) []GearSet {
	var gearSets []GearSet
	m := g.ToMap()

	gearMeldCombos := map[string][]GearItem{}

	for slot := range m {
		gearMeldCombos[slot] = GearMeldCombinations(materiaTypes, m[slot]...)

	}

	stop := false
	for !stop {
		stop = true

		gsMap := map[string]GearItem{}
		for slot, items := range gearMeldCombos {
			if len(items) > 0 {
				stop = false
				gsMap[slot] = items[0]
				gearMeldCombos[slot] = items[1:]
			}
		}
		gs := GearSet{}
		gs.LoadFromMap(gsMap)
		gearSets = append(gearSets, gs)
	}

	return gearSets
}

func (g AvailableGear) BiS(job Job, lvl Level, clan Clan, materiaTypes []*Materia) GearSet {
	gearSets := g.Combinations(materiaTypes)

	bestDmg := math.Inf(-1)
	var bis GearSet
	for _, gs := range gearSets {

		gcd := GCD(lvl, gs.Stats().SKS)
		if gcd > 2.5 {
			continue
		}

		gs.Job = job
		gs.Lvl = lvl
		gs.Clan = clan
		dmg := gs.DamageNormalized()

		if dmg > bestDmg {
			bis = gs
			bestDmg = dmg
		}
	}

	return bis
}
