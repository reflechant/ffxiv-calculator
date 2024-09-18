package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMateriaCombinationsOf1(t *testing.T) {
	combos := MateriaCombinations(MateriaTypes, 1)
	cnt := 0
	for range combos {
		cnt++
	}
	assert.Equal(t, len(MateriaTypes), cnt)
}

func TestMateriaCombinationsOf2(t *testing.T) {
	combos := MateriaCombinations([]*Materia{
		SavageAim11,
		SavageAim12,
	}, 2)

	assert.ElementsMatch(t, [][]*Materia{
		{},
		{SavageAim11},
		{SavageAim12},
		{SavageAim11, SavageAim11},
		{SavageAim12, SavageAim12},
		{SavageAim11, SavageAim12},
		{SavageAim12, SavageAim11},
	}, slices.Collect(combos))
}

func TestGearMeldCombinations(t *testing.T) {
	gear := Gear.Item("Skyruin Gunblade")
	gearMeldCombos := slices.Collect(GearMeldCombinations([]*Materia{
		SavageAim11,
		SavageAim12,
	}, gear))

	assert.ElementsMatch(t, []GearItem{
		gear,
		gear.Meld(SavageAim11),
		gear.Meld(SavageAim12),
		gear.Meld(SavageAim11).Meld(SavageAim11),
		gear.Meld(SavageAim12).Meld(SavageAim12),
		gear.Meld(SavageAim11).Meld(SavageAim12),
		gear.Meld(SavageAim12).Meld(SavageAim11),
	}, gearMeldCombos)
}
