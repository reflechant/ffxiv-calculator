package main

import (
	"fmt"
	"iter"
	"maps"
	"math"
)

type AvailableGear struct {
	Weapon  []GearItem
	OffHand []GearItem

	Head  []GearItem
	Body  []GearItem
	Hands []GearItem
	Legs  []GearItem
	Feet  []GearItem

	Ears  []GearItem
	Neck  []GearItem
	Wrist []GearItem
	Ring  []GearItem
}

var SlotItemUICategory = map[string]string{
	"Weapon":  "Weapon",
	"OffHand": "OffHand",
	"Head":    "Head",
	"Body":    "Body",
	"Hands":   "Hands",
	"Legs":    "Legs",
	"Feet":    "Feet",
	"Ears":    "Earrings",
	"Neck":    "Necklace",
	"Wrist":   "Bracelets",
	"Ring":    "Ring",
}

func (ag AvailableGear) Map() map[string][]GearItem {
	return map[string][]GearItem{
		"Weapon":  ag.Weapon,
		"OffHand": ag.OffHand,
		"Head":    ag.Head,
		"Body":    ag.Body,
		"Hands":   ag.Hands,
		"Legs":    ag.Legs,
		"Feet":    ag.Feet,
		"Ears":    ag.Ears,
		"Neck":    ag.Neck,
		"Wrist":   ag.Wrist,
		"Ring":    ag.Ring,
	}
}

func (ag AvailableGear) SlotAvailableItems() map[string][]GearItem {
	return map[string][]GearItem{
		"Weapon":    ag.Weapon,
		"OffHand":   ag.OffHand,
		"Head":      ag.Head,
		"Body":      ag.Body,
		"Hands":     ag.Hands,
		"Legs":      ag.Legs,
		"Feet":      ag.Feet,
		"Ears":      ag.Ears,
		"Neck":      ag.Neck,
		"Wrist":     ag.Wrist,
		"LeftRing":  ag.Ring,
		"RightRing": ag.Ring,
	}
}

func (ag *AvailableGear) Load(m map[string][]GearItem) {
	ag.Weapon = m["Weapon"]
	ag.OffHand = m["OffHand"]
	ag.Head = m["Head"]
	ag.Body = m["Body"]
	ag.Hands = m["Hands"]
	ag.Legs = m["Legs"]
	ag.Feet = m["Feet"]
	ag.Ears = m["Ears"]
	ag.Neck = m["Neck"]
	ag.Wrist = m["Wrist"]
	ag.Ring = m["Ring"]
}

// Combinations will return all possible gear set combinations out of available gear
func (g AvailableGear) Combinations(materiaTypes []*Materia) iter.Seq[GearSet] {
	return func(yield func(GearSet) bool) {
		slotItems := g.SlotAvailableItems()
		slots := make([]string, 0, len(slotItems))
		gearMeldCombos := map[string]iter.Seq[GearItem]{}

		fmt.Printf("slotItems: %v\n", slotItems)

		for slot, items := range slotItems {
			slots = append(slots, slot)
			gearMeldCombos[slot] = GearMeldCombinations(materiaTypes, items...)
		}

		type queueItem struct {
			gsMap       map[string]GearItem
			slotsToFill []string
		}

		queue := []queueItem{
			{
				gsMap:       map[string]GearItem{},
				slotsToFill: slots,
			},
		}

		fmt.Printf("queue: %v\n", queue)

		for len(queue) > 0 {
			gsMap, slotsToFill := queue[0].gsMap, queue[0].slotsToFill
			queue = queue[1:]

			if len(slotsToFill) == 0 {
				gs := GearSet{}
				gs.LoadFromMap(gsMap)
				// skip invalid gear sets with duplicate unique items (it can only be rings)
				if (gs.LeftRing.Name == gs.RightRing.Name) && gs.LeftRing.IsUnique {
					// fmt.Printf("can't have 2 of %s\n", gs.LeftRing.Name)
					continue
				}

				if !yield(gs) {
					return
				}

				continue
			}

			slot := slotsToFill[0]
			items := gearMeldCombos[slot]
			for it := range items {
				m := maps.Clone(gsMap)
				m[slot] = it
				queue = append(queue, queueItem{
					gsMap:       m,
					slotsToFill: slotsToFill[1:],
				})
			}
		}
	}
}

func (g AvailableGear) BiS(job Job, lvl Level, clan Clan, materiaTypes []*Materia, gcdMin, gcdMax float64) GearSet {
	// gearSets := slices.Collect(g.Combinations(materiaTypes))
	gearSets := g.Combinations(materiaTypes)

	fmt.Println(gearSets)

	bestDmg := math.Inf(-1)
	var bis GearSet
	for gs := range gearSets {

		gcd := GCD(lvl, job, gs.Stats().SecondaryStats)
		if gcd < gcdMin || gcd > gcdMax {
			continue
		}

		gs.Job = job
		gs.Lvl = lvl
		gs.Clan = clan
		dmg := gs.DamageNormalized()

		if dmg > bestDmg {
			// fmt.Printf("better gear set found (dmg = %f):\n", dmg)
			// diffJSON, _ := json.MarshalIndent(bis.Stats().Diff(gs.Stats()), "", "  ")
			// fmt.Println(string(diffJSON))

			bis = gs
			bestDmg = dmg
		}
	}

	return bis
}
