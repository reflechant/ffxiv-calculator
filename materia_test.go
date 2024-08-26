package main

import (
	"fmt"
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

	for i := range combos {
		for j := range combos[i] {
			fmt.Printf("%v ", combos[i][j].Name)
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
