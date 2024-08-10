package main

import (
	"embed"
	"encoding/json"
)

type GearSet struct {
	Job     Job
	Weapon  GearItem
	OffHand GearItem

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
	return DamageBase(Attributes{
		Lvl: LevelMod100,
		Job: set.Job,
		WD:  int(set.Weapon.WD()), // it's always integer, it being float is an artifact of data scraping
		AP:  set.Stats().STR,
	}, 100)
}

func (set GearSet) DamageNormalized() float64 {
	return DamageNormalized(Attributes{
		Lvl: LevelMod100,
		Job: set.Job,
		WD:  int(set.Weapon.WD()), // it's always integer, it being float is an artifact of data scraping
		AP:  set.Stats().STR,
	}, 100)
}

// GearItem is a generic peace of gear. Some fields are type-dependant and only contain non-zero values for certain types of gear
type GearItem struct {
	Name   string `json:"name"`
	Lvl    uint   `json:"ilvl"`
	Jobs   Job    // bitmask
	JobLvl uint   `json:"job level"`
	Stats
	PhysDMG       float64 `json:"Physical Damage"`
	MagDMG        float64 `json:"Magic Damage"`
	AutoAtk       float64 `json:"Auto-attack"`
	Delay         float64 `json:"Delay"`
	MateriaSlots  int     `json:"materia slots"`
	MateriaMelded []Materia
}

func (it GearItem) EffectiveStats() Stats {
	cap := it.SecondaryStatCap()

	secondaryStats := it.Stats.SecondaryStats
	for _, m := range it.MateriaMelded {
		secondaryStats = SumSecondaryStats(secondaryStats, m.SecondaryStats)
	}

	return Stats{
		MainStats:      it.Stats.MainStats,
		SecondaryStats: secondaryStats.Cap(cap),
	}
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
	return GearItem{
		Stats:         it.Stats,
		MateriaSlots:  it.MateriaSlots + 1,
		MateriaMelded: append(it.MateriaMelded, materia),
	}
}

func (it GearItem) WD() float64 {
	if it.PhysDMG == 0.0 {
		return it.MagDMG
	}

	return it.PhysDMG
}

//go:embed items.json
var f embed.FS

func LoadGear() ([]GearItem, error) {
	data, _ := f.ReadFile("items.json")

	var gear []GearItem
	err := json.Unmarshal(data, &gear)

	return gear, err
}
