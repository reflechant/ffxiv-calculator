package main

// All the formulas and values come from the Allagan Studies project
// (see https://www.akhmorning.com/)

func BaseStats(lvl Level, job Job, clan Clan) Stats {
	return Stats{
		MainStats: MainStats{
			STR: lvl.Main()*job.Stats().STR/100 + clan.Stats().STR,
			DEX: lvl.Main()*job.Stats().DEX/100 + clan.Stats().DEX,
			VIT: lvl.Main()*job.Stats().VIT/100 + clan.Stats().VIT,
			INT: lvl.Main()*job.Stats().INT/100 + clan.Stats().INT,
			MND: lvl.Main()*job.Stats().MND/100 + clan.Stats().MND,
		},
		SecondaryStats: SecondaryStats{
			DET:  lvl.Main(),
			PT:   lvl.Main(),
			CRIT: lvl.Sub(),
			DH:   lvl.Sub(),
			SKS:  lvl.Sub(),
			SPS:  lvl.Sub(),
			TNC:  lvl.Sub(),
		},
	}
}

func SpeedFactor(lvl Level, SS int) int {
	// uses skill speed or spell speed
	return 130 * (SS - lvl.Sub()) / lvl.Div()
}

func GCD(lvl Level, job Job, stats SecondaryStats) float64 {
	return (1000 - float64(SpeedFactor(lvl, job.SS(stats)))) * float64(2.5) / 1000
}

func AttackFactor(lvl Level, AP int, job Job) int {
	// uses attack power (AP) or magic attack potency (MAP)
	if (job & TANK) > 0 {
		return 190*(AP-lvl.Main())/lvl.Main() + 100
	}

	return 237*(AP-lvl.Main())/lvl.Main() + 100
}

func DeterminationFactor(lvl Level, DET int) int {
	return 140*(DET-lvl.Main())/lvl.Div() + 1000
}

func TenacityFactor(lvl Level, TNC int) int {
	return 112*(TNC-lvl.Sub())/lvl.Div() + 1000
}

func WeaponDamageFactor(lvl Level, job Job, WD int) int {
	// WD is Physical Damage or Magical Damage of the weapon
	// the attribute being used depends on the action, we assume it's the primary stat for simplicity
	return lvl.Main()*job.PrimaryStat(job.Stats().MainStats)/1000 + WD
}

// Direct Hit (DH damage is 1.25x of normal)

const dhMultiplier float64 = 1.25

func DirectHitChance(lvl Level, DH int) float64 {
	return float64(uint(550*(DH-lvl.Sub())/lvl.Div())) / 10
}

// Crit

func CritChance(lvl Level, CRIT int) float64 {
	return float64(uint(200*(CRIT-lvl.Sub())/lvl.Div()+50)) / 10
}

func CritMultiplier(lvl Level, CRIT int) float64 {
	return float64(200*(CRIT-lvl.Sub())/lvl.Div()+1400) / 1000
}

// Damage (DMG)

type Attributes struct {
	Lvl  Level
	Job  Job
	WD   int // PhysDmg or MagDmg
	AP   int // Attack Power or Magic Attack Potency
	DET  int
	TNC  int
	CRIT int
	DH   int
}

// DamageBase returns damage without CRIT or DH. Damage randomization (+- 5%) not accounted for.
func DamageBase(attr Attributes, potency int) int {
	D1 := potency * AttackFactor(attr.Lvl, attr.AP, attr.Job) / 100
	D2 := D1 * DeterminationFactor(attr.Lvl, attr.DET) / 1000
	D3 := (D2 * TenacityFactor(attr.Lvl, attr.TNC)) / 1000
	D4 := D3 * WeaponDamageFactor(attr.Lvl, attr.Job, attr.WD) / 100
	return D4
}

// DamageNormalized returns average damage accounting for CRIT and DH chance. Damage randomization (+- 5%) not accounted for.
func DamageNormalized(attr Attributes, potency int) float64 {
	D := DamageBase(attr, potency)
	D2 := normalize(float64(D), CritMultiplier(attr.Lvl, attr.CRIT), CritChance(attr.Lvl, attr.CRIT))
	D3 := normalize(D2, dhMultiplier, DirectHitChance(attr.Lvl, attr.DH))
	return D3
}

// normalize returns normalized base value according to multiplier and its probability(chance). Formula taken from etro.gg's help page. Multiplier is expected to be >= 1.
func normalize(base, multiplier, chance float64) float64 {
	return base * (1 + chance/100*(multiplier-1))
}
