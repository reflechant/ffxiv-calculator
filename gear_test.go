package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		// "one materia type, one slot -> two combinations (empty and melded)": {
		// 	slotsNum:     1,
		// 	materiaTypes: []Materia{SavageAim9},
		// 	expected: [][]Materia{
		// 		[]Materia{},
		// 		[]Materia{SavageAim9},
		// 	},
		// },
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

func TestGearLoader(t *testing.T) {
	jsonStr := `
	{
        "type": "Legs",
        "name": "Light-heavy Brais of Casting",
        "ilvl": 710,
        "job": [
            "THM",
            "ACN",
            "BLM",
            "SMN",
            "RDM",
            "BLU",
            "PCT"
        ],
        "job level": 100,
        "Defense": 589.0,
        "Magic Defense": 1031.0,
        "Vitality": 494,
        "Intelligence": 530,
        "Determination": 357,
        "Direct Hit Rate": 250,
        "materia slots": 2
    }`

	g := GearItem{}
	expected := GearItem{
		Name:   "Light-heavy Brais of Casting",
		Lvl:    710,
		JobLvl: 100,
		Stats: Stats{
			MainStats: MainStats{
				VIT: 494,
				INT: 530,
			},
			SecondaryStats: SecondaryStats{
				DET: 357,
				DH:  250,
			},
		},
		MateriaSlots: 2,
	}

	err := json.Unmarshal([]byte(jsonStr), &g)
	require.NoError(t, err)

	assert.Equal(t, expected, g)
}
