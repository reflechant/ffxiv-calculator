package main

type Materia struct {
	Name string
	SecondaryStats
}

var (
	SavageAim9  = Materia{Name: "Savage Aim Materia IX", SecondaryStats: SecondaryStats{CRIT: 12}}
	SavageAim10 = Materia{Name: "Savage Aim Materia X", SecondaryStats: SecondaryStats{CRIT: 36}}
	SavageAim11 = Materia{Name: "Savage Aim Materia XI", SecondaryStats: SecondaryStats{CRIT: 18}}
	SavageAim12 = Materia{Name: "Savage Aim Materia XII", SecondaryStats: SecondaryStats{CRIT: 54}}

	SavageMight9  = Materia{Name: "Savage Might Materia IX", SecondaryStats: SecondaryStats{DET: 12}}
	SavageMight10 = Materia{Name: "Savage Might Materia X", SecondaryStats: SecondaryStats{DET: 36}}
	SavageMight11 = Materia{Name: "Savage Might Materia XI", SecondaryStats: SecondaryStats{DET: 18}}
	SavageMight12 = Materia{Name: "Savage Might Materia XII", SecondaryStats: SecondaryStats{DET: 54}}

	HeavensEye9  = Materia{Name: "Heavens' Eye Materia IX", SecondaryStats: SecondaryStats{DH: 12}}
	HeavensEye10 = Materia{Name: "Heavens' Eye Materia X", SecondaryStats: SecondaryStats{DH: 36}}
	HeavensEye11 = Materia{Name: "Heavens' Eye Materia XI", SecondaryStats: SecondaryStats{DH: 18}}
	HeavensEye12 = Materia{Name: "Heavens' Eye Materia XII", SecondaryStats: SecondaryStats{DH: 54}}

	QuickArm9  = Materia{Name: "Quick Arm Materia IX", SecondaryStats: SecondaryStats{SKS: 12}}
	QuickArm10 = Materia{Name: "Quick Arm Materia X", SecondaryStats: SecondaryStats{SKS: 36}}
	QuickArm11 = Materia{Name: "Quick Arm Materia XI", SecondaryStats: SecondaryStats{SKS: 18}}
	QuickArm12 = Materia{Name: "Quick Arm Materia XII", SecondaryStats: SecondaryStats{SKS: 54}}

	QuickTongue9  = Materia{Name: "Quick Tongue Materia IX", SecondaryStats: SecondaryStats{SPS: 12}}
	QuickTongue10 = Materia{Name: "Quick Tongue Materia X", SecondaryStats: SecondaryStats{SPS: 36}}
	QuickTongue11 = Materia{Name: "Quick Tongue Materia XI", SecondaryStats: SecondaryStats{SPS: 18}}
	QuickTongue12 = Materia{Name: "Quick Tongue Materia XII", SecondaryStats: SecondaryStats{SPS: 54}}

	Battledance9  = Materia{Name: "Battledance Materia IX", SecondaryStats: SecondaryStats{TNC: 12}}
	Battledance10 = Materia{Name: "Battledance Materia X", SecondaryStats: SecondaryStats{TNC: 36}}
	Battledance11 = Materia{Name: "Battledance Materia XI", SecondaryStats: SecondaryStats{TNC: 18}}
	Battledance12 = Materia{Name: "Battledance Materia XII", SecondaryStats: SecondaryStats{TNC: 54}}

	Piety9  = Materia{Name: "Piety Materia IX", SecondaryStats: SecondaryStats{PT: 12}}
	Piety10 = Materia{Name: "Piety Materia X", SecondaryStats: SecondaryStats{PT: 36}}
	Piety11 = Materia{Name: "Piety Materia XI", SecondaryStats: SecondaryStats{PT: 18}}
	Piety12 = Materia{Name: "Piety Materia XII", SecondaryStats: SecondaryStats{PT: 54}}
)
