package main

// All the formulas and values come from the Allagan Studies project
// (see https://www.akhmorning.com/)

func Speed(lvl LevelMod, SS uint) uint {
	// uses skill speed or spell speed
	return 130*(SS-lvl.Sub)/lvl.Div + 1000
}

func AttackFactor(lvl LevelMod, AP uint) uint {
	// uses attack power (AP) or magic attack potency (MAP)
	// tank=190, others=237
	return 190*(AP-lvl.Main)/lvl.Main + 100
}

func DeterminationFactor(lvl LevelMod, DET uint) uint {
	return 140*(DET-lvl.Main)/lvl.Div + 1000
}

func TenacityFactor(lvl LevelMod, TNC uint) uint {
	return 112*(TNC-lvl.Sub)/lvl.Div + 1000
}

func WeaponDamageFactor(lvl LevelMod, job JobMod, WD uint) uint {
	// weaponDmg is Physical Damage or Magical Damage of the weapon
	// the attribute being used depends on the action, (STR for GNB)
	return lvl.Main*uint(job.STR)/1000 + WD
}

// Direct Hit (DH damage is 1.25x of normal)

const dhMultiplier float64 = 1.25

func DirectHitChance(lvl LevelMod, DH uint) float64 {
	return float64(uint(550*(DH-lvl.Sub)/lvl.Div)) / 10
}

// Crit

func CritChance(lvl LevelMod, CRIT uint) float64 {
	return float64(uint(200*(CRIT-lvl.Sub)/lvl.Div+50)) / 10
}

func CritMultiplier(lvl LevelMod, CRIT uint) float64 {
	return float64(200*(CRIT-lvl.Sub)/lvl.Div+1400) / 1000
}

// GCD

func GCD() float64 {
	panic("not implemented yet")
}

// Damage (DMG)

type Attributes struct {
	Lvl  LevelMod
	Job  JobMod
	WD   uint
	AP   uint
	DET  uint
	TNC  uint
	CRIT uint
	DH   uint
}

// DamageBase returns damage without CRIT or DH. Damage randomization (+- 5%) not accounted for.
func DamageBase(attr Attributes, potency uint) uint {
	D1 := potency * AttackFactor(attr.Lvl, attr.AP) / 100
	D2 := D1 * DeterminationFactor(attr.Lvl, attr.DET) / 1000
	D3 := (D2 * TenacityFactor(attr.Lvl, attr.TNC)) / 1000
	D4 := D3 * WeaponDamageFactor(attr.Lvl, attr.Job, attr.WD) / 100
	return D4
}

// DamageNormalized returns average damage accounting for CRIT and DH chance. Damage randomization (+- 5%) not accounted for.
func DamageNormalized(attr Attributes, potency uint) float64 {
	D := DamageBase(attr, potency)
	D2 := normalize(float64(D), CritMultiplier(attr.Lvl, attr.CRIT), CritChance(attr.Lvl, attr.CRIT))
	D3 := normalize(D2, dhMultiplier, DirectHitChance(attr.Lvl, attr.DH))
	return D3
}

// normalize returns normalized base value according to multiplier and its probability(chance). Formula taken from etro.gg's help page. Multiplier is expected to be >= 1.
func normalize(base, multiplier, chance float64) float64 {
	return base * (1 + chance/100*(multiplier-1))
}

// Attributes
// Strength

func STR(lvl LevelMod, job JobMod, clan ClanMod) uint {
	return uint(lvl.Main*uint(job.STR)/100) + uint(job.STR)
}
