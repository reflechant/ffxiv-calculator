package main

import (
	"fmt"
)

func main() {
	gearMap := GearMap()

	set := GearSet{
		Lvl:       Lvl100,
		Job:       GNB,
		Clan:      KeepersOfTheMoon,
		Weapon:    gearMap.Item("Skyruin Gunblade").Meld(SavageMight12).Meld(SavageMight12),
		Head:      gearMap.Item("Light-heavy Bandana of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Body:      gearMap.Item("Archeo Kingdom Cuirass of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Hands:     gearMap.Item("Light-heavy Halfgloves of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Legs:      gearMap.Item("Archeo Kingdom Poleyns of Fending").Meld(SavageAim12).Meld(SavageAim12),
		Feet:      gearMap.Item("Light-heavy Greaves of Fending").Meld(HeavensEye12).Meld(HeavensEye12),
		Ears:      gearMap.Item("Archeo Kingdom Earrings of Fending").Meld(SavageAim12).Meld(HeavensEye12),
		Neck:      gearMap.Item("Dark Horse Champion's Choker of Fending"),
		Wrist:     gearMap.Item("Light-heavy Bangle of Fending").Meld(HeavensEye12),
		LeftRing:  gearMap.Item("Light-heavy Ring of Fending").Meld(HeavensEye12),
		RightRing: gearMap.Item("Archeo Kingdom Ring of Fending").Meld(HeavensEye12),
	}

	// statsJSON, err := json.MarshalIndent(set.Stats(), "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Gear set stats: %v\n", string(statsJSON))

	fmt.Printf("damage base: %v\n", set.DamageBase())
	// fmt.Printf("damage normalized: %v\n", set.DamageNormalized())
}
