package main

import (
	"log"
	"slices"

	"github.com/charmbracelet/huh"
)

func runForm() (AvailableGear, []*Materia) {
	gear := AvailableGear{}

	// gearMap := gear.ToMap()
	selectedGearMap := map[string]*[]string{}
	// slotSelects := []huh.Field{}
	slotGroups := []*huh.Group{}

	for slot := range slices.Values([]string{"Weapon",
		"OffHand",
		"Head",
		"Body",
		"Hands",
		"Legs",
		"Feet",
		"Ears",
		"Neck",
		"Wrist",
		"LeftRing",
		"RightRing"}) {
		selectedItems := []string{}

		items := Gear.Index["GNB"][SlotItemUICategory[slot]]
		huhOptions := []huh.Option[string]{}
		for it := range slices.Values(items) {
			huhOptions = append(huhOptions, huh.NewOption(it, it))
		}
		selectedGearMap[slot] = &selectedItems

		if len(huhOptions) == 0 {
			continue
		}

		// slotSelects = append(slotSelects, huh.NewMultiSelect[string]().Title(slot).Options(huhOptions...).Value(&selectedItems))
		slotGroups = append(slotGroups,
			huh.NewGroup(
				huh.NewMultiSelect[string]().Options(huhOptions...).Value(&selectedItems),
			).Title(slot).WithWidth(50),
		)

	}

	// materiaTypeIndex := map[string]*Materia{}
	// materiaTypeNames := []string{}
	// for m := range slices.Values(MateriaTypes) {
	// 	materiaTypeIndex[m.Name] = m
	// 	materiaTypeNames = append(materiaTypeNames, m.Name)
	// }
	// materiaOptions := []huh.Option[string]{}
	// for name := range slices.Values(materiaTypeNames) {
	// 	materiaOptions = append(materiaOptions, huh.NewOption[string](name, name))
	// }
	// selectedMateriaTypes := []string{}

	// construct the form

	// slotGroups = append(slotGroups,
	// 	huh.NewGroup(
	// 		huh.NewMultiSelect[string]().Title("available materia").Options(materiaOptions...).Value(&selectedMateriaTypes),
	// 	),
	// )
	form := huh.NewForm(
		slotGroups...,
	// huh.NewGroup(
	// add select job
	// slotSelects...,
	// ),
	// huh.NewGroup(
	// 	huh.NewMultiSelect[string]().Title("available materia").Options(materiaOptions...).Value(&selectedMateriaTypes),
	// ),
	).WithLayout(huh.LayoutGrid(4, 3))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// calculate selected materia types
	materiaTypes := []*Materia{}
	// for s := range slices.Values(selectedMateriaTypes) {
	// 	if s, ok := materiaTypeIndex[s]; ok {
	// 		materiaTypes = append(materiaTypes, s)
	// 	}
	// }

	// calculate selected gear
	gearMap := gear.ToMap()
	for slot, itemNames := range selectedGearMap {
		items := []GearItem{}
		for itemName := range slices.Values(*itemNames) {
			items = append(items, Gear.Item(itemName))
		}
		gearMap[slot] = items
	}
	gear.LoadFromMap(gearMap)

	return gear, materiaTypes
}
