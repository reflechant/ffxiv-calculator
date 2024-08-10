package main

const actionDelay = 2500 // milliseconds

type ClanMod MainStats

var ClanModifiers = map[string]ClanMod{
	"Keepers of the Moon": {STR: -1, DEX: 2, VIT: -2, INT: 1, MND: 3},
}
