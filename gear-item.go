package main

import (
	"fmt"
	"strings"
)

// GearItem is a generic peace of gear. Some fields are type-dependant and only contain non-zero values for certain types of gear
type GearItem struct {
	Name         string
	Type         string
	EquipSlot    int     `json:"EquipSlotCategory"`
	Lvl          uint    `json:"ItemLvl"`
	JobLvl       uint    `json:"EquipLvl"`
	PhysDMG      float64 `json:"PhysDmg,omitempty"`
	MagDMG       float64 `json:"MagDmg,omitempty"`
	Delay        float64 `json:"DelayMS,omitempty"`
	AutoAtk      float64 `json:"Auto-attack,omitempty"`
	MateriaSlots int     `json:"MateriaSlotCount,omitempty"`
	CanBeHQ      bool    `json:"CanBeHq,omitempty"`
	IsUnique     bool    `json:"IsUnique,omitempty"`
	Stats
	BaseParamSpecial Stats
	MateriaMelded    []*Materia
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

func (it *GearItem) MakeHQ() {
	if !it.CanBeHQ {
		return
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

func (it GearItem) Meld(materia *Materia) GearItem {
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

func (it GearItem) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%50s", it.Name))
	for _, m := range it.MateriaMelded {
		b.WriteString("  ")
		b.WriteString(m.String())
	}

	return b.String()
}
