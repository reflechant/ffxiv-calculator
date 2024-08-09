package main

// All the formulas and values come from the Allagan Studies project
// (see https://www.akhmorning.com/)

func Speed(lvl LevelMod, SS uint) uint {
	// uses skill speed or spell speed
	return 130*(SS-lvl.Sub)/lvl.Div + 1000
}

func TankAttack(AP uint) uint {
	// uses attack power (AP) or magic attack potency (MAP)
	return uint(115*(AP-340)/340) + 100
}

func Determination(lvl LevelMod, DET uint) uint {
	return 130*(DET-lvl.Main)/lvl.Div + 1000
}

func Tenacity(lvl LevelMod, TNC uint) uint {
	return 100*(TNC-lvl.Sub)/lvl.Div + 1000
}

func WeaponDamage(lvl LevelMod, job JobMod, weaponDmg uint) uint {
	// weaponDmg is Physical Damage or Magical Damage of the weapon
	// the attribute being used depends on the action, (STR for GNB)
	return lvl.Main*uint(job.STR)/1000 + weaponDmg
}

// Direct Hit (DH damage is 1.25x of normal)

func DirectHitChance(lvl LevelMod, DH uint) float32 {
	return float32(uint(550*(DH-lvl.Sub)/lvl.Div)) / 10
}

// Crit

func CritChance(lvl LevelMod, CRIT uint) float32 {
	return float32(uint(200*(CRIT-lvl.Sub)/lvl.Div+50)) / 10
}

func CritDmg(lvl LevelMod, CRIT uint) uint {
	return 200*(CRIT-lvl.Sub)/lvl.Div + 1400
}

// GCD

func GCD() float32 {
	return 0.0
}

// Damage (DMG)

func DamageBeforeCritDH(lvl LevelMod, job JobMod, potency uint, WD, AP, DET, TNC uint) float64 {
	D1 := ((potency * TankAttack(AP) * Determination(lvl, DET)) / 100) / 1000
	D2 := ((D1 * Tenacity(lvl, TNC)) / 1000) * WeaponDamage(lvl, job, WD) / 100
	D3 := D2 // crit and dh modifiers

	return float64(D3) //* math.Sqrt(0.01/12) // +-5% damage variance, uniform distribution
}

// Attributes
// Strength

func STR(lvl LevelMod, job JobMod, clan ClanMod) uint {
	return uint(lvl.Main*uint(job.STR)/100) + uint(job.STR)
}
