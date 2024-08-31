package main

type Job uint64

const (
	// tanks
	GLA Job = 1 << iota
	PLD
	MRD
	WAR
	DRK
	GNB

	// healers
	CNJ
	WHM
	SCH
	AST
	SGE

	// melee DPS
	LNC
	DRG
	PGL
	MNK
	ROG
	NIN
	SAM
	RPR
	VPR

	// ranged physical DPS
	ARC
	BRD
	MCH
	DNC

	// ranged magical DPS
	THM
	BLM
	ACN
	SMN
	RDM
	PCT

	// limited
	BLU

	// roles
	TANK   Job = GLA | PLD | MRD | WAR | DRK | GNB
	HEALER Job = CNJ | WHM | SCH | AST | SGE

	MELEE_DPS           Job = LNC | DRG | PGL | MNK | ROG | NIN | SAM | RPR | VPR
	RANGED_PHYSICAL_DPS Job = ARC | BRD | MCH | DNC
	RANGED_MAGICAL_DPS  Job = THM | BLM | ACN | SMN | RDM | PCT
	DPS                 Job = MELEE_DPS | RANGED_PHYSICAL_DPS | RANGED_MAGICAL_DPS
)

type JobMod struct {
	HP int
	MP int
	MainStats
}

var jobModifiers = map[Job]JobMod{
	GNB: {HP: 120, MP: 100, MainStats: MainStats{STR: 100, VIT: 110, DEX: 95, INT: 60, MND: 100}},
	BRD: {HP: 105, MP: 100, MainStats: MainStats{STR: 90, VIT: 100, DEX: 115, INT: 85, MND: 80}},
}

func (job Job) Stats() JobMod {
	return jobModifiers[job]
}

// stat cap multipliers in percents
func (job Job) StatCapMultipliers() Stats {
	statsMultipliers := Stats{
		MainStats: MainStats{
			STR: 100,
			DEX: 100,
			VIT: 100,
			INT: 100,
			MND: 100,
		},
		SecondaryStats: SecondaryStats{
			CRIT: 100,
			DET:  100,
			DH:   100,
			SKS:  100,
			SPS:  100,
			TNC:  100,
			PT:   100,
		},
	}
	if job&(HEALER|RANGED_MAGICAL_DPS) > 0 {
		statsMultipliers.VIT = 90
	}

	return statsMultipliers
}

func (job Job) PrimaryStat(stats MainStats) int {
	if job&TANK > 0 {
		return stats.STR
	}
	if job&RANGED_PHYSICAL_DPS > 0 {
		return stats.DEX
	}
	if job&(ROG|NIN|VPR) > 0 {
		return stats.DEX
	}
	if job&MELEE_DPS > 0 {
		return stats.STR
	}
	if job&RANGED_MAGICAL_DPS > 0 {
		return stats.INT
	}

	// RDM, AST, SGE ???

	panic("not implemented")

}

// SS returns skill speed or spell speed
func (job Job) SS(stats SecondaryStats) int {
	if job&(TANK|MELEE_DPS|RANGED_PHYSICAL_DPS) > 0 {
		return stats.SKS
	}

	return stats.SPS
}

var mainArmCategories = map[Job][]string{
	// tanks
	GLA: {"Gladiator's Arm"},
	PLD: {"Gladiator's Arm"},
	MRD: {"Marauder's Arm"},
	WAR: {"Marauder's Arm"},
	DRK: {"Dark Knight's Arm"},
	GNB: {"Gunbreaker's Arm"},

	// healers
	CNJ: {"One–handed Conjurer's Arm", "Two–handed Conjurer's Arm"},
	WHM: {"One–handed Conjurer's Arm", "Two–handed Conjurer's Arm"},
	SCH: {"Scholar's Arm"},
	AST: {"Astrologian's Arm"},
	SGE: {"Sage's Arm"},

	// melee DPS
	LNC: {"Lancer's Arm"},
	DRG: {"Lancer's Arm"},
	PGL: {"Pugilist's Arm"},
	MNK: {"Pugilist's Arm"},
	ROG: {"Rogue's Arm"},
	NIN: {"Rogue's Arm"},
	SAM: {"Samurai's Arm"},
	RPR: {"Reaper's Arm"},
	VPR: {"Viper's Arm"},

	// ranged physical DPS
	ARC: {"Archer's Arm"},
	BRD: {"Archer's Arm"},
	MCH: {"Machinist's Arm"},
	DNC: {"Dancer's Arm"},

	// ranged magical DPS
	THM: {"One–handed Thaumaturge's Arm", "Two–handed Thaumaturge's Arm"},
	BLM: {"One–handed Thaumaturge's Arm", "Two–handed Thaumaturge's Arm"},
	ACN: {"Arcanist's Grimoire"},
	SMN: {"Arcanist's Grimoire"},
	RDM: {"Red Mage's Arm"},
	PCT: {"Pictomancer's Arm"},

	// limited
	BLU: {"Blue Mage's Arm"},
}

func (job Job) mainArmCategories() []string {
	return mainArmCategories[job]
}

func (job Job) String() string {
	return jobs[job]
}

var jobs = map[Job]string{
	// tanks
	// GLA: "GLA ",
	PLD: "PLD",
	// MRD: "MRD",
	WAR: "WAR",
	DRK: "DRK",
	GNB: "GNB",

	// healers
	// CNJ: "CNJ",
	WHM: "WHM",
	SCH: "SCH",
	AST: "AST",
	SGE: "SGE",

	// melee DPS
	// LNC: "LNC",
	DRG: "DRG",
	// PGL: "PGL",
	MNK: "MNK",
	// ROG: "ROG",
	NIN: "NIN",
	SAM: "SAM",
	RPR: "RPR",
	VPR: "VPR",

	// ranged physical DPS
	// ARC: "ARC",
	BRD: "BRD",
	MCH: "MCH",
	DNC: "DNC",

	// ranged magical DPS
	// THM: "THM",
	BLM: "BLM",
	// ACN: "ACN",
	SMN: "SMN",
	RDM: "RDM",
	PCT: "PCT",

	// limited
	BLU: "BLU",
}
