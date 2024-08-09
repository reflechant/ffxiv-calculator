package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareBaseDamageWithXIVGear(t *testing.T) {
	// This test compares damage per 100 potency with XIVGear.app using this gearset:
	// https://xivgear.app/?page=sl%7Cdab85390-9b1c-4c89-a745-90ff45ff39c6

	dmg := DamageBase(Attributes{
		Lvl:  LevelMod100,
		Job:  JobModifiers["GNB"],
		WD:   141,
		AP:   4395,
		DET:  1978,
		TNC:  794,
		CRIT: 3006,
		DH:   1068,
	}, 100)

	// this test fails but I expect this values to be at least close which is not the case
	// 3669 is XIVGear.app calculated DMG/100potency for the gearset above without Crit and DH
	assert.InEpsilon(t, uint(3656), dmg, 0.001)
}

func TestCompareNormalisedDamageWithEtro(t *testing.T) {
	dmgNormalized := DamageNormalized(Attributes{
		Lvl:  LevelMod100,
		Job:  JobModifiers["GNB"],
		WD:   141,
		AP:   4395,
		DET:  1978,
		TNC:  794,
		CRIT: 3006,
		DH:   1068,
	}, 100)
	assert.InEpsilon(t, 4292.86, dmgNormalized, 0.001)
}

func TestCritChance(t *testing.T) {
	chance := CritChance(LevelMod100, 3006)
	// 23.6% taken from xivgear.app
	assert.Equal(t, 23.6, chance)
}

func TestCritMulti(t *testing.T) {
	chance := CritMultiplier(LevelMod100, 3006)
	// 1.586 taken from xivgear.app
	assert.Equal(t, 1.586, chance)
}

func TestDHChance(t *testing.T) {
	chance := DirectHitChance(LevelMod100, 1068)
	// 12.8% taken from xivgear.app
	assert.Equal(t, 12.8, chance)
}

func TestATKMultiplier(t *testing.T) {
	atk := AttackFactor(LevelMod100, 4395)
	assert.Equal(t, 18.07, float64(atk)/100)
}

func TestDeterminationMultiplier(t *testing.T) {
	det := DeterminationFactor(LevelMod100, 1978)
	assert.Equal(t, 1.077, float64(det)/1000)
}
