package main

import (
	"fmt"
	"reflect"
	"strings"
)

type GearSet struct {
	Lvl  Level
	Job  Job
	Clan Clan

	Weapon  GearItem `json:"Weapon,omitempty"`
	OffHand GearItem `json:"OffHand,omitempty"`

	Head  GearItem `json:"Head,omitempty"`
	Body  GearItem `json:"Body,omitempty"`
	Hands GearItem `json:"Hands,omitempty"`
	Legs  GearItem `json:"Legs,omitempty"`
	Feet  GearItem `json:"Feet,omitempty"`

	Ears      GearItem `json:"Ears,omitempty"`
	Neck      GearItem `json:"Neck,omitempty"`
	Wrist     GearItem `json:"Wrist,omitempty"`
	LeftRing  GearItem `json:"LeftRing,omitempty"`
	RightRing GearItem `json:"RightRing,omitempty"`
}

func (gs GearSet) Map() map[string]GearItem {
	return map[string]GearItem{
		"Weapon":    gs.Weapon,
		"OffHand":   gs.OffHand,
		"Head":      gs.Head,
		"Body":      gs.Body,
		"Hands":     gs.Hands,
		"Legs":      gs.Legs,
		"Feet":      gs.Feet,
		"Ears":      gs.Ears,
		"Neck":      gs.Neck,
		"Wrist":     gs.Wrist,
		"LeftRing":  gs.LeftRing,
		"RightRing": gs.RightRing,
	}
}

func (gs *GearSet) LoadFromMap(m map[string]GearItem) {
	gs.Weapon = m["Weapon"]
	gs.OffHand = m["OffHand"]
	gs.Head = m["Head"]
	gs.Body = m["Body"]
	gs.Hands = m["Hands"]
	gs.Legs = m["Legs"]
	gs.Feet = m["Feet"]
	gs.Ears = m["Ears"]
	gs.Neck = m["Neck"]
	gs.Wrist = m["Wrist"]
	gs.LeftRing = m["LeftRing"]
	gs.RightRing = m["RightRing"]
}

func (set GearSet) Stats() Stats {
	return SumStats(
		BaseStats(set.Lvl, set.Job, set.Clan),
		set.Weapon.EffectiveStats(),
		set.OffHand.EffectiveStats(),

		set.Head.EffectiveStats(),
		set.Body.EffectiveStats(),
		set.Hands.EffectiveStats(),
		set.Legs.EffectiveStats(),
		set.Feet.EffectiveStats(),

		set.Ears.EffectiveStats(),
		set.Neck.EffectiveStats(),
		set.Wrist.EffectiveStats(),
		set.LeftRing.EffectiveStats(),
		set.RightRing.EffectiveStats(),
	)
}

func (set GearSet) DamageBase() int {
	stats := set.Stats()

	return DamageBase(Attributes{
		Lvl:  set.Lvl,
		Job:  set.Job,
		WD:   int(set.Weapon.WD()), // it's always integer, it being float is an artifact of data scraping
		AP:   set.Job.PrimaryStat(stats.MainStats),
		DET:  stats.DET,
		TNC:  stats.TNC,
		CRIT: stats.CRIT,
		DH:   stats.DH,
	}, 100)
}

func (set GearSet) DamageNormalized() float64 {
	stats := set.Stats()

	return DamageNormalized(Attributes{
		Lvl:  set.Lvl,
		Job:  set.Job,
		WD:   int(set.Weapon.WD()), // it's always integer, it being float is an artifact of data scraping
		AP:   set.Job.PrimaryStat(stats.MainStats),
		DET:  stats.DET,
		TNC:  stats.TNC,
		CRIT: stats.CRIT,
		DH:   stats.DH,
	}, 100)
}

func (set GearSet) String() string {
	var b strings.Builder
	b.WriteString("gear set:\n")
	b.WriteString(fmt.Sprintf("%s lvl %d, %s\n", jobs[set.Job], set.Lvl, set.Clan))
	b.WriteString(fmt.Sprintf("GCD: %.02f\n", GCD(set.Lvl, set.Job, set.Stats().SecondaryStats)))
	// m := set.Map()

	setV := reflect.ValueOf(set)
	for i := 3; i < setV.NumField(); i++ {
		slot := setV.Type().Field(i).Name
		b.WriteString(fmt.Sprintf("%10s -> ", slot))
		item := setV.Field(i).Interface().(GearItem)
		if item.Name != "" {
			b.WriteString(item.String())
			// b.WriteString(fmt.Sprintf("%10s -> %s\n", slot, item))
		}
		b.WriteByte('\n')
	}

	return b.String()
}
