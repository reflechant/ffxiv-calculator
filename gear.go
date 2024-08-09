package main

// Item is a generic peace of gear
type Item struct {
	STR int
	DEX int
	VIT int
	INT int
	MND int

	CRIT uint
	DET  uint
	DH   uint

	TNC uint
	PT  uint

	MateriaSlots int
}

type GearSet struct {
	Weapon    Item
	OffWeapon Item
	Head      Item
	Body      Item
	Hands     Item
	Legs      Item
	Feet      Item
	Ears      Item
	Neck      Item
	Wrist     Item
	LeftRing  Item
	RightRing Item
}
