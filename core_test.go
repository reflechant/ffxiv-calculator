package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const epsilon = 0.002

func TestCompareBaseDamageWithXIVGear(t *testing.T) {
	// This test compares damage per 100 potency with XIVGear.app using this gearset:
	// https://xivgear.app/?page=sl%7Cdab85390-9b1c-4c89-a745-90ff45ff39c6

	dmg := DamageBase(Attributes{
		Lvl:  100,
		Job:  GNB,
		WD:   141,
		AP:   4395,
		DET:  1978,
		TNC:  794,
		CRIT: 3006,
		DH:   1068,
	}, 100)
	assert.InEpsilon(t, uint(3656), dmg, epsilon)
}

func TestCompareNormalisedDamageWithEtro(t *testing.T) {
	dmgNormalized := DamageNormalized(Attributes{
		Lvl:  100,
		Job:  GNB,
		WD:   141,
		AP:   4395,
		DET:  1978,
		TNC:  794,
		CRIT: 3006,
		DH:   1068,
	}, 100)
	assert.InEpsilon(t, 4292.86, dmgNormalized, epsilon)
}

func TestCritChance(t *testing.T) {
	chance := CritChance(100, 3006)
	// 23.6% taken from xivgear.app
	assert.Equal(t, 23.6, chance)
}

func TestCritMulti(t *testing.T) {
	chance := CritMultiplier(100, 3006)
	// 1.586 taken from xivgear.app
	assert.InEpsilon(t, 1.586, chance, epsilon)
}

func TestDHChance(t *testing.T) {
	chance := DirectHitChance(100, 1068)
	// 12.8% taken from xivgear.app
	assert.InEpsilon(t, 12.8, chance, epsilon)
}

func TestATKMultiplier(t *testing.T) {
	atk := AttackFactor(100, 4395, GNB)
	assert.Equal(t, 18.07, float64(atk)/100)
}

func TestDeterminationMultiplier(t *testing.T) {
	det := DeterminationFactor(100, 1978)
	assert.Equal(t, 1.077, float64(det)/1000)
}

func TestBaseStats(t *testing.T) {
	stats := BaseStats(100, GNB, KeepersOfTheMoon)
	assert.Equal(t, 439, stats.STR)
}

func TestGCD(t *testing.T) {
	gcd := GCD(100, 702)
	assert.InEpsilon(t, 2.46, gcd, 0.0031)
}
