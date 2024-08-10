package main

type Materia struct {
	Name string
	SecondaryStats
}

var (
	SavageAimIX  = Materia{Name: "Savage Aim Materia IX", SecondaryStats: SecondaryStats{CRIT: 12}}
	SavageAimX   = Materia{Name: "Savage Aim Materia X", SecondaryStats: SecondaryStats{CRIT: 36}}
	SavageAimXI  = Materia{Name: "Savage Aim Materia XI", SecondaryStats: SecondaryStats{CRIT: 18}}
	SavageAimXII = Materia{Name: "Savage Aim Materia XII", SecondaryStats: SecondaryStats{CRIT: 54}}

	SavageMightIX  = Materia{Name: "Savage Might Materia IX", SecondaryStats: SecondaryStats{DET: 12}}
	SavageMightX   = Materia{Name: "Savage Might Materia X", SecondaryStats: SecondaryStats{DET: 36}}
	SavageMightXI  = Materia{Name: "Savage Might Materia XI", SecondaryStats: SecondaryStats{DET: 18}}
	SavageMightXII = Materia{Name: "Savage Might Materia XII", SecondaryStats: SecondaryStats{DET: 54}}

	HeavensEyeIX  = Materia{Name: "Heavens' Eye Materia IX", SecondaryStats: SecondaryStats{DH: 12}}
	HeavensEyeX   = Materia{Name: "Heavens' Eye Materia X", SecondaryStats: SecondaryStats{DH: 36}}
	HeavensEyeXI  = Materia{Name: "Heavens' Eye Materia XI", SecondaryStats: SecondaryStats{DH: 18}}
	HeavensEyeXII = Materia{Name: "Heavens' Eye Materia XII", SecondaryStats: SecondaryStats{DH: 54}}

	QuickArmIX  = Materia{Name: "Quick Arm Materia IX", SecondaryStats: SecondaryStats{SKS: 12}}
	QuickArmX   = Materia{Name: "Quick Arm Materia X", SecondaryStats: SecondaryStats{SKS: 36}}
	QuickArmXI  = Materia{Name: "Quick Arm Materia XI", SecondaryStats: SecondaryStats{SKS: 18}}
	QuickArmXII = Materia{Name: "Quick Arm Materia XII", SecondaryStats: SecondaryStats{SKS: 54}}

	QuickTongueIX  = Materia{Name: "Quick Tongue Materia IX", SecondaryStats: SecondaryStats{SPS: 12}}
	QuickTongueX   = Materia{Name: "Quick Tongue Materia X", SecondaryStats: SecondaryStats{SPS: 36}}
	QuickTongueXI  = Materia{Name: "Quick Tongue Materia XI", SecondaryStats: SecondaryStats{SPS: 18}}
	QuickTongueXII = Materia{Name: "Quick Tongue Materia XII", SecondaryStats: SecondaryStats{SPS: 54}}

	BattledanceIX  = Materia{Name: "Battledance Materia IX", SecondaryStats: SecondaryStats{TNC: 12}}
	BattledanceX   = Materia{Name: "Battledance Materia X", SecondaryStats: SecondaryStats{TNC: 36}}
	BattledanceXI  = Materia{Name: "Battledance Materia XI", SecondaryStats: SecondaryStats{TNC: 18}}
	BattledanceXII = Materia{Name: "Battledance Materia XII", SecondaryStats: SecondaryStats{TNC: 54}}

	PietyIX  = Materia{Name: "Piety Materia IX", SecondaryStats: SecondaryStats{PT: 12}}
	PietyX   = Materia{Name: "Piety Materia X", SecondaryStats: SecondaryStats{PT: 36}}
	PietyXI  = Materia{Name: "Piety Materia XI", SecondaryStats: SecondaryStats{PT: 18}}
	PietyXII = Materia{Name: "Piety Materia XII", SecondaryStats: SecondaryStats{PT: 54}}
)
