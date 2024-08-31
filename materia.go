package main

import (
	"fmt"
	"iter"
	"reflect"
	"slices"
)

type Materia struct {
	Name string
	SecondaryStats
}

var (
	SavageAim9  = &Materia{Name: "Savage Aim Materia IX", SecondaryStats: SecondaryStats{CRIT: 12}}
	SavageAim10 = &Materia{Name: "Savage Aim Materia X", SecondaryStats: SecondaryStats{CRIT: 36}}
	SavageAim11 = &Materia{Name: "Savage Aim Materia XI", SecondaryStats: SecondaryStats{CRIT: 18}}
	SavageAim12 = &Materia{Name: "Savage Aim Materia XII", SecondaryStats: SecondaryStats{CRIT: 54}}

	SavageMight9  = &Materia{Name: "Savage Might Materia IX", SecondaryStats: SecondaryStats{DET: 12}}
	SavageMight10 = &Materia{Name: "Savage Might Materia X", SecondaryStats: SecondaryStats{DET: 36}}
	SavageMight11 = &Materia{Name: "Savage Might Materia XI", SecondaryStats: SecondaryStats{DET: 18}}
	SavageMight12 = &Materia{Name: "Savage Might Materia XII", SecondaryStats: SecondaryStats{DET: 54}}

	HeavensEye9  = &Materia{Name: "Heavens' Eye Materia IX", SecondaryStats: SecondaryStats{DH: 12}}
	HeavensEye10 = &Materia{Name: "Heavens' Eye Materia X", SecondaryStats: SecondaryStats{DH: 36}}
	HeavensEye11 = &Materia{Name: "Heavens' Eye Materia XI", SecondaryStats: SecondaryStats{DH: 18}}
	HeavensEye12 = &Materia{Name: "Heavens' Eye Materia XII", SecondaryStats: SecondaryStats{DH: 54}}

	QuickArm9  = &Materia{Name: "Quick Arm Materia IX", SecondaryStats: SecondaryStats{SKS: 12}}
	QuickArm10 = &Materia{Name: "Quick Arm Materia X", SecondaryStats: SecondaryStats{SKS: 36}}
	QuickArm11 = &Materia{Name: "Quick Arm Materia XI", SecondaryStats: SecondaryStats{SKS: 18}}
	QuickArm12 = &Materia{Name: "Quick Arm Materia XII", SecondaryStats: SecondaryStats{SKS: 54}}

	QuickTongue9  = &Materia{Name: "Quick Tongue Materia IX", SecondaryStats: SecondaryStats{SPS: 12}}
	QuickTongue10 = &Materia{Name: "Quick Tongue Materia X", SecondaryStats: SecondaryStats{SPS: 36}}
	QuickTongue11 = &Materia{Name: "Quick Tongue Materia XI", SecondaryStats: SecondaryStats{SPS: 18}}
	QuickTongue12 = &Materia{Name: "Quick Tongue Materia XII", SecondaryStats: SecondaryStats{SPS: 54}}

	Battledance9  = &Materia{Name: "Battledance Materia IX", SecondaryStats: SecondaryStats{TNC: 12}}
	Battledance10 = &Materia{Name: "Battledance Materia X", SecondaryStats: SecondaryStats{TNC: 36}}
	Battledance11 = &Materia{Name: "Battledance Materia XI", SecondaryStats: SecondaryStats{TNC: 18}}
	Battledance12 = &Materia{Name: "Battledance Materia XII", SecondaryStats: SecondaryStats{TNC: 54}}

	Piety9  = &Materia{Name: "Piety Materia IX", SecondaryStats: SecondaryStats{PT: 12}}
	Piety10 = &Materia{Name: "Piety Materia X", SecondaryStats: SecondaryStats{PT: 36}}
	Piety11 = &Materia{Name: "Piety Materia XI", SecondaryStats: SecondaryStats{PT: 18}}
	Piety12 = &Materia{Name: "Piety Materia XII", SecondaryStats: SecondaryStats{PT: 54}}
)

var MateriaTypes = []*Materia{
	SavageAim9,
	SavageAim10,
	SavageAim11,
	SavageAim12,
	SavageMight9,
	SavageMight10,
	SavageMight11,
	SavageMight12,
	HeavensEye9,
	HeavensEye10,
	HeavensEye11,
	HeavensEye12,
	QuickArm9,
	QuickArm10,
	QuickArm11,
	QuickArm12,
	QuickTongue9,
	QuickTongue10,
	QuickTongue11,
	QuickTongue12,
	Battledance9,
	Battledance10,
	Battledance11,
	Battledance12,
	Piety9,
	Piety10,
	Piety11,
	Piety12,
}

func (m *Materia) String() string {
	mStats := reflect.ValueOf(m.SecondaryStats)
	for i := 0; i < mStats.NumField(); i++ {
		statName := mStats.Type().Field(i).Name
		statVal, ok := mStats.Field(i).Interface().(int)
		if !ok {
			panic("non int stat")
		}
		if statVal > 0 {
			return fmt.Sprintf("%4s+%2d", statName, statVal)
		}
	}

	return "??? materia"

	// statsJSON, err := json.Marshal(m.SecondaryStats)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return string(statsJSON)
}

// MateriaCombinations returns combinations with replacement
func MateriaCombinations(materiaTypes []*Materia, slotsNum int) iter.Seq[[]*Materia] {
	// result := [][]*Materia{{}}
	// combination := make([]*Materia, slots)

	// var generate func(int, int)
	// generate = func(index, start int) {
	// 	if index == slots {
	// 		temp := make([]*Materia, slots)
	// 		copy(temp, combination)
	// 		result = append(result, temp)
	// 		return
	// 	}

	// 	for i := start; i < len(materiaTypes); i++ {
	// 		combination[index] = materiaTypes[i]
	// 		generate(index+1, i)
	// 	}

	// 	if index > 0 {
	// 		temp := make([]*Materia, index)
	// 		copy(temp, combination[:index])
	// 		result = append(result, temp)
	// 	}
	// }

	// generate(0, 0)

	return func(yield func([]*Materia) bool) {
		var generate func([]*Materia)
		generate = func(m []*Materia) {
			if len(m) == slotsNum {
				if !yield(m) {
					return
				}
			}
			for mType := range slices.Values(materiaTypes) {
				generate(append(slices.Clone(m), mType))
			}
		}
		generate([]*Materia{})
	}
}

func GearMeldCombinations(materiaTypes []*Materia, items ...GearItem) []GearItem {
	result := []GearItem{{}}

	for _, item := range items {
		materiaCombos := MateriaCombinations(materiaTypes, item.MateriaSlots)
		for combo := range materiaCombos {
			g := item
			for _, materia := range combo {
				g = g.Meld(materia)
			}
			result = append(result, g)
		}
	}

	return result
}
