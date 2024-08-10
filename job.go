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
}

func (job Job) Stats() JobMod {
	return jobModifiers[job]
}

func (job Job) PrimaryStat(stats MainStats) int {
	switch job {
	case GNB:
		return stats.STR
	default:
		panic("not implemented")
	}
}

// SS returns skill speed or spell speed
func (job Job) SS() int {
	panic("not implemented")
}
