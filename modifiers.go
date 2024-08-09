package main

type LevelMod struct {
	HP   uint
	MP   uint
	Main uint
	Sub  uint
	Div  uint
}

var LevelMod100 = LevelMod{
	HP:   4000,
	MP:   10000,
	Main: 440,
	Sub:  420,
	Div:  2780,
}

const actionDelay = 2500 // milliseconds

type ClanMod struct {
	STR int
	DEX int
	VIT int
	INT int
	MND int
}

var ClanModifiers = map[string]ClanMod{
	"Keepers of the Moon": {STR: -1, DEX: 2, VIT: -2, INT: 1, MND: 3},
}

type JobMod struct {
	HP  int
	MP  int
	STR int
	DEX int
	VIT int
	INT int
	MND int
}

var JobModifiers = map[string]JobMod{
	"GNB": {HP: 120, MP: 100, STR: 100, VIT: 110, DEX: 95, INT: 60, MND: 100},
}
