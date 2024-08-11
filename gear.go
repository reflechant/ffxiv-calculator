package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
)

type GearSet struct {
	Lvl  Level
	Job  Job
	Clan Clan

	Weapon GearItem
	Shield GearItem

	Head  GearItem
	Body  GearItem
	Hands GearItem
	Legs  GearItem
	Feet  GearItem

	Ears      GearItem
	Neck      GearItem
	Wrist     GearItem
	LeftRing  GearItem
	RightRing GearItem
}

func (set GearSet) Stats() Stats {
	return SumStats(
		BaseStats(set.Lvl, set.Job, set.Clan),
		set.Weapon.EffectiveStats(),
		set.Shield.EffectiveStats(),

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
		Lvl:  Lvl100,
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
		Lvl:  Lvl100,
		Job:  set.Job,
		WD:   int(set.Weapon.WD()), // it's always integer, it being float is an artifact of data scraping
		AP:   set.Job.PrimaryStat(stats.MainStats),
		DET:  stats.DET,
		TNC:  stats.TNC,
		CRIT: stats.CRIT,
		DH:   stats.DH,
	}, 100)
}

// GearItem is a generic peace of gear. Some fields are type-dependant and only contain non-zero values for certain types of gear
type GearItem struct {
	Name string `json:"name"`
	Lvl  uint   `json:"ilvl"`
	// Jobs   Job    // bitmask
	JobLvl uint `json:"job level"`
	Stats
	PhysDMG float64 `json:"Physical Damage,omitempty"`
	MagDMG  float64 `json:"Magic Damage,omitempty"`
	// AutoAtk       float64 `json:"Auto-attack,omitempty"`
	// Delay         float64 `json:"Delay,omitempty"`
	MateriaSlots  int `json:"materia slots,omitempty"`
	MateriaMelded []Materia
}

func (it GearItem) EffectiveStats() Stats {
	cap := it.SecondaryStatCap()

	secondaryStats := it.Stats.SecondaryStats
	for _, m := range it.MateriaMelded {
		secondaryStats = SumSecondaryStats(secondaryStats, m.SecondaryStats)
	}

	stats := Stats{
		MainStats:      it.Stats.MainStats,
		SecondaryStats: secondaryStats.Cap(cap),
	}

	return stats
}

func (it GearItem) SecondaryStatCap() int {
	st := it.Stats
	return max(st.CRIT, st.DET, st.DH, st.SKS, st.SPS, st.TNC, st.PT)
}

// materiaTypes indicate possible types of materia, we assume you have infinite amount of these
func (it GearItem) PossibleMelds(materiaTypes []Materia) [][]Materia {
	variants := make([][]Materia, 0)
	return variants
}

func (it GearItem) Meld(materia Materia) GearItem {
	if len(it.MateriaMelded) >= it.MateriaSlots {
		return it
	}
	it.MateriaSlots++
	it.MateriaMelded = append(it.MateriaMelded, materia)

	return it
}

func (it GearItem) WD() float64 {
	return max(it.PhysDMG, it.MagDMG)
}

// JSONs produced by scraping Eorzea Database (with eorzea_spider.py) miss some items (for example on August 11, 2024 Resilient gear was still hidden and marked with ??? (probably to avoid spoilers?))

//go:embed items.json
var f embed.FS

// job -> item type -> item name -> item stats
type GearDB map[string]map[string]map[string]GearItem

func LoadGearJSON() GearDB {
	data, err := f.ReadFile("items.json")
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

// possible categories:
// "weapon",
// "Shield",
// "Head",
// "Body",
// "Legs",
// "Hands",
// "Feet",
// "Necklace"
// "Earrings"
// "Bracelets
// "Ring",

func (db GearDB) Item(name string) GearItem {
	for job := range db {
		for category := range db[job] {
			v, ok := db[job][category][name]
			if ok {
				fmt.Printf("found %s in %s: %s\n", name, job, category)
				return v
			}
		}
	}

	panic("not found " + name)
}
