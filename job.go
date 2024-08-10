package main

type Job uint64

const (
	// tanks
	GLA = 1 << iota
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
)

type JobMod struct {
	HP int
	MP int
	MainStats
}

var JobModifiers = map[Job]JobMod{
	GNB: {HP: 120, MP: 100, MainStats: MainStats{STR: 100, VIT: 110, DEX: 95, INT: 60, MND: 100}},
}

func (job Job) PrimaryStat() int {
	switch job {
	case GNB:
		return JobModifiers[GNB].STR
	default:
		return JobModifiers[GNB].STR
	}
}

// SS returns skill speed or spell speed
func (job Job) SS() int {
	switch job {
	case GNB:
		return JobModifiers[GNB].STR
	default:
		return JobModifiers[GNB].STR
	}
}
