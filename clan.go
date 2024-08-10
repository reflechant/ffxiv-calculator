package main

const (
	// Hyur
	Midlanders  Clan = "Midlanders"
	Highlanders Clan = "Highlanders"
	// Elezen
	Wildwood  Clan = "Wildwood"
	Duskwight Clan = "Duskwight"
	// Lalafell
	Plainsfolk Clan = "Plainsfolk"
	Dunesfolk  Clan = "Dunesfolk"
	// Miqo'te
	SeekersOfTheSun  Clan = "Seekers of the Sun"
	KeepersOfTheMoon Clan = "Keepers of the Moon"
	// Roegadyn
	SeaWolves  Clan = "Sea Wolves"
	Hellsguard Clan = "Hellsguard"
	// Au Ra
	Raen  Clan = "Raen"
	Xaela Clan = "Xaela"
	// Viera
	Rava  Clan = "Rava"
	Veena Clan = "Veena"
	// Hrothgar
	Helions Clan = "Helions"
	TheLost Clan = "The Lost"
)

var clanModifiers = map[Clan]MainStats{
	KeepersOfTheMoon: {STR: -1, DEX: 2, VIT: -2, INT: 1, MND: 3},
}

type Clan string

func (clan Clan) Stats() MainStats {
	return clanModifiers[clan]
}
