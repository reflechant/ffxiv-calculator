package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCombinations(t *testing.T) {
// 	t.Parallel()

// 	input := AvailableGear{
// 		Weapon: []GearItem{
// 			Gear.Item("Skyruin Gunblade"),
// 			Gear.Item("Dark Horse Champion's Gunblade"),
// 		},
// 		LeftRing: []GearItem{
// 			Gear.Item("Resilient Ring of Fending"),
// 			Gear.Item("Light-heavy Ring of Fending"),
// 		},
// 	}

// 	items := input.Combinations([]*Materia{})

// 	assert.ElementsMatch(t, []GearItem{
// 		Gear.Item("Skyruin Gunblade"),
// 		Gear.Item("Dark Horse Champion's Gunblade"),
// 		Gear.Item("Resilient Ring of Fending"),
// 		Gear.Item("Light-heavy Ring of Fending"),
// 	}, items)
// }

func TestBis(t *testing.T) {
	inventory := AvailableGear{
		Weapon: []GearItem{
			Gear.Item("Skyruin Gunblade"),
		},
	}

	bis := inventory.BiS(GNB, 100, KeepersOfTheMoon, []*Materia{SavageAim12, SavageMight12, HeavensEye12, QuickArm12})

	assert.Equal(t, GearSet{
		Lvl:    100,
		Job:    GNB,
		Clan:   KeepersOfTheMoon,
		Weapon: Gear.Item("Skyruin Gunblade").Meld(SavageMight12).Meld(SavageMight12),
	}, bis)
}
