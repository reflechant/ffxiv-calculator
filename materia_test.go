package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMateriaCombinationsOf1(t *testing.T) {
	combos := MateriaCombinations(MateriaTypes, 1)
	assert.Len(t, combos, len(MateriaTypes)+1)
}

func TestMateriaCombinationsOf2(t *testing.T) {
	combos := MateriaCombinations([]*Materia{
		SavageAim11,
		SavageAim12,
	}, 2)

	for combo := range combos {
		for m := range slices.Values(combo) {
			fmt.Printf("%v ", m.Name)
		}
		fmt.Println()
	}

	assert.ElementsMatch(t, [][]*Materia{
		{},
		{SavageAim11},
		{SavageAim12},
		{SavageAim11, SavageAim11},
		{SavageAim12, SavageAim12},
		{SavageAim11, SavageAim12},
	}, combos)
}
