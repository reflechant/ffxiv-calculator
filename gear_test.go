package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGNBGearSet(t *testing.T) {
	t.Skip()
}

func TestGearItem_PossibleMelds(t *testing.T) {
	tests := map[string]struct {
		slotsNum     int
		materiaTypes []Materia
		expected     [][]Materia
	}{
		"empty materia list -> no possible melds": {
			slotsNum:     2,
			materiaTypes: []Materia{},
			expected:     [][]Materia{},
		},
		"one materia type, one slot -> two combinations (empty and melded)": {
			slotsNum:     1,
			materiaTypes: []Materia{SavageAimIX},
			expected: [][]Materia{
				[]Materia{},
				[]Materia{SavageAimIX},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			it := GearItem{
				MateriaSlots: tt.slotsNum,
			}

			possibleMelds := it.PossibleMelds(tt.materiaTypes)
			assert.Equal(t, tt.expected, possibleMelds)
		})
	}
}
