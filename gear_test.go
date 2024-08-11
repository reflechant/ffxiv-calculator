package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var skyruinGunblade = GearItem{
	Name:    "Skyruin Gunblade",
	Lvl:     710,
	Jobs:    GNB,
	JobLvl:  100,
	PhysDMG: 141.0,
	AutoAtk: 131.6,
	Delay:   2.8,
	Stats: Stats{
		MainStats: MainStats{
			STR: 550,
			VIT: 570,
		},
		SecondaryStats: SecondaryStats{
			CRIT: 370,
			DET:  259,
		},
	},
	MateriaSlots: 2,
}

func TestGearSet_Stats_NoGear100(t *testing.T) {
	set := GearSet{
		Lvl:  Lvl100,
		Job:  GNB,
		Clan: KeepersOfTheMoon,
	}
	stats := set.Stats()
	assert.Equal(t, 439, stats.STR)
	assert.Equal(t, 420, stats.CRIT)
	assert.Equal(t, 420, stats.DH)
	assert.Equal(t, 440, stats.DET)
	assert.Equal(t, 420, stats.TNC)
}

func TestGearSet_Stats_OnlyMainArm(t *testing.T) {
	set := GearSet{
		Lvl:    Lvl100,
		Job:    GNB,
		Clan:   KeepersOfTheMoon,
		Weapon: skyruinGunblade,
	}
	stats := set.Stats()
	assert.Equal(t, 989, stats.STR)
	assert.Equal(t, 790, stats.CRIT)
	assert.Equal(t, 420, stats.DH)
	assert.Equal(t, 699, stats.DET)
	assert.Equal(t, 420, stats.TNC)
}

func TestGearSet_DamageNormalized_OnlyMainArm(t *testing.T) {
	set := GearSet{
		Lvl:    Lvl100,
		Job:    GNB,
		Clan:   KeepersOfTheMoon,
		Weapon: skyruinGunblade,
	}
	dmg := set.DamageNormalized()
	assert.InEpsilon(t, 651.429256, dmg, 0.002)
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

func TestGearSet_Stats_GNBMelded(t *testing.T) {
	gearMap := GearMapFromJSON()
	set := GearSet{
		Lvl:       Lvl100,
		Job:       GNB,
		Clan:      KeepersOfTheMoon,
		Weapon:    gearMap.Item("Skyruin Gunblade").Meld(SavageMight12).Meld(SavageMight12),
		Head:      gearMap.Item("Light-heavy Bandana of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Body:      gearMap.Item("Archeo Kingdom Cuirass of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Hands:     gearMap.Item("Light-heavy Halfgloves of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Legs:      gearMap.Item("Archeo Kingdom Poleyns of Fending").Meld(SavageAim12).Meld(SavageAim12),
		Feet:      gearMap.Item("Light-heavy Greaves of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Ears:      gearMap.Item("Archeo Kingdom Earrings of Fending").Meld(SavageAim12).Meld(HeavensEye12),
		Neck:      gearMap.Item("Dark Horse Champion's Choker of Fending"),
		Wrist:     gearMap.Item("Light-heavy Bangle of Fending").Meld(HeavensEye12),
		LeftRing:  gearMap.Item("Light-heavy Ring of Fending").Meld(HeavensEye12),
		RightRing: gearMap.Item("Archeo Kingdom Ring of Fending").Meld(HeavensEye12),
	}

	stats := set.Stats()
	assert.Equal(t, 4395, stats.STR)
	assert.Equal(t, 2829, stats.CRIT)
	assert.Equal(t, 1068, stats.DH)
	assert.Equal(t, 2102, stats.DET)
	assert.Equal(t, 847, stats.TNC)
}

func TestGearSet_Stats_BardMelded(t *testing.T) {
	gearMap := GearMapFromJSON()

	set := GearSet{
		Lvl:       Lvl100,
		Job:       BRD,
		Clan:      KeepersOfTheMoon,
		Weapon:    gearMap.Item("Gastraphetes"),
		Head:      gearMap.Item("Light-heavy Mask of Aiming"),
		Body:      gearMap.Item("Light-heavy Jacket of Aiming"),
		Hands:     gearMap.Item("Light-heavy Gloves of Aiming"),
		Legs:      gearMap.Item("Light-heavy Brais of Aiming"),
		Feet:      gearMap.Item("Light-heavy Boots of Aiming"),
		Ears:      gearMap.Item("Light-heavy Earring of Aiming"),
		Neck:      gearMap.Item("Light-heavy Choker of Aiming"),
		Wrist:     gearMap.Item("Light-heavy Bangle of Aiming"),
		LeftRing:  gearMap.Item("Light-heavy Ring of Aiming").Meld(SavageAim12),
		RightRing: gearMap.Item("Archeo Kingdom Ring of Aiming").Meld(HeavensEye12),
	}

	stats := set.Stats()
	assert.Equal(t, 4378, stats.DEX)
	assert.Equal(t, 1889, stats.CRIT)
	assert.Equal(t, 1664, stats.DH)
	assert.Equal(t, 1613, stats.DET)

	dmg := set.DamageNormalized()
	assert.InEpsilon(t, 6039.99569, dmg, epsilon)
	// expected dmg: 5277; 6039.99569
}

func TestGearSet_Stats_NoMateria(t *testing.T) {
	gearMap := GearMapFromJSON()

	set := GearSet{
		Lvl:       Lvl100,
		Job:       GNB,
		Clan:      KeepersOfTheMoon,
		Weapon:    gearMap.Item("Skyruin Gunblade"),
		Head:      gearMap.Item("Light-heavy Bandana of Fending"),
		Body:      gearMap.Item("Archeo Kingdom Cuirass of Fending"),
		Hands:     gearMap.Item("Light-heavy Halfgloves of Fending"),
		Legs:      gearMap.Item("Archeo Kingdom Poleyns of Fending"),
		Feet:      gearMap.Item("Light-heavy Greaves of Fending"),
		Ears:      gearMap.Item("Archeo Kingdom Earrings of Fending"),
		Neck:      gearMap.Item("Dark Horse Champion's Choker of Fending"),
		Wrist:     gearMap.Item("Light-heavy Bangle of Fending"),
		LeftRing:  gearMap.Item("Light-heavy Ring of Fending"),
		RightRing: gearMap.Item("Archeo Kingdom Ring of Fending"),
	}

	stats := set.Stats()
	assert.Equal(t, 4395, stats.STR)
	assert.Equal(t, 2669, stats.CRIT)
	assert.Equal(t, 420, stats.DH)
	assert.Equal(t, 1994, stats.DET)
	assert.Equal(t, 847, stats.TNC)

	damageBase := set.DamageBase()
	assert.InEpsilon(t, 3666, damageBase, 0.002)
}
