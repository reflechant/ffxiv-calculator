package main

// GearItem is a generic peace of gear
type GearItem struct {
	baseStats     Stats
	MateriaSlots  int
	MateriaMelded []Materia
}

func (it GearItem) Stats() Stats {
	cap := it.SecondaryStatCap()

	secondaryStats := it.baseStats.SecondaryStats
	for _, m := range it.MateriaMelded {
		secondaryStats = SumSecondaryStats(secondaryStats, m.SecondaryStats)
	}

	return Stats{
		MainStats:      it.baseStats.MainStats,
		SecondaryStats: secondaryStats.Cap(cap),
	}
}

func (it GearItem) SecondaryStatCap() uint {
	st := it.baseStats
	return max(st.CRIT, st.DET, st.DH, st.SKS, st.SPS, st.TNC, st.PT)
}

// materiaTypes indicate possible types of materia, we assume you have infinite amount of these
func (it GearItem) PossibleMelds(materiaTypes []Materia) [][]Materia {
	variants := make([][]Materia, 0)
	return variants
}

func (it GearItem) Meld(materia Materia) (GearItem, bool) {
	if len(it.MateriaMelded) >= it.MateriaSlots {
		return it, false
	}
	return GearItem{
		baseStats:     it.baseStats,
		MateriaSlots:  it.MateriaSlots + 1,
		MateriaMelded: append(it.MateriaMelded, materia),
	}, true
}

type GearSet struct {
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
		set.Weapon.Stats(),
		set.OffHand.Stats(),

		set.Head.Stats(),
		set.Body.Stats(),
		set.Hands.Stats(),
		set.Legs.Stats(),
		set.Feet.Stats(),

		set.Ears.Stats(),
		set.Neck.Stats(),
		set.Wrist.Stats(),
		set.LeftRing.Stats(),
		set.RightRing.Stats(),
	)
}
