package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareDmgWithXIVGear(t *testing.T) {
	// This test compares damage per 100 potency with XIVGear.app using this gearset:
	//https://xivgear.app/?page=sl%7Cd4c6cbf2-c3fe-4d7e-843a-b84116d0dd87

	dmg := DamageBeforeCritDH(LevelMod100, JobModifiers["GNB"], 100, 141, 4395, 1978, 967)

	// this test fails but I expect this values to be at least close which is not the case
	// 3669 is XIVGear.app calculated DMG/100potency for the gearset above without Crit and DH
	assert.Equal(t, int(3669), int(dmg))
}
