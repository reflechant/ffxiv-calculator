package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	gearMap := GearMap()

	set := GearSet{
		Lvl:       Lvl100,
		Job:       GNB,
		Clan:      KeepersOfTheMoon,
		Weapon:    gearMap["Skyruin Gunblade"].Meld(SavageMight12).Meld(SavageMight12),
		Head:      gearMap["Light-heavy Bandana of Fending"].Meld(HeavensEye12).Meld(HeavensEye12),
		Body:      gearMap["Archeo Kingdom Cuirass of Fending"].Meld(HeavensEye12).Meld(HeavensEye12),
		Hands:     gearMap["Light-heavy Halfgloves of Fending"].Meld(HeavensEye12).Meld(HeavensEye12),
		Legs:      gearMap["Archeo Kingdom Poleyns of Fending"].Meld(SavageAim12).Meld(SavageAim12),
		Feet:      gearMap["Light-heavy Greaves of Fending"].Meld(HeavensEye12).Meld(HeavensEye12),
		Ears:      gearMap["Archeo Kingdom Earrings of Fending"].Meld(SavageAim12).Meld(HeavensEye12),
		Neck:      gearMap["Dark Horse Champion's Choker of Fending"],
		Wrist:     gearMap["Light-heavy Bangle of Fending"].Meld(HeavensEye12),
		LeftRing:  gearMap["Resilient Ring of Fending"].Meld(HeavensEye12),
		RightRing: gearMap["Archeo Kingdom Ring of Fending"].Meld(HeavensEye12),
	}

	statsJSON, err := json.MarshalIndent(set.Stats(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gear set stats: %v\n", string(statsJSON))

	fmt.Printf("damage base: %v\n", set.DamageBase())
	// fmt.Printf("damage normalized: %v\n", set.DamageNormalized())
}
