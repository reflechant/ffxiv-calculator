package main

type LevelMod struct {
	HP   int
	MP   int
	Main int
	Sub  int
	Div  int
}

var LevelMod100 = LevelMod{
	HP:   4000,
	MP:   10000,
	Main: 440,
	Sub:  420,
	Div:  2780,
}
