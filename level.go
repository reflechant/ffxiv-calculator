package main

type Level int

type LevelMod struct {
	HP   int
	MP   int
	Main int
	Sub  int
	Div  int
}

var levelMods = map[Level]LevelMod{
	1:   {HP: 86, MP: 10000, Main: 20, Sub: 56, Div: 56},
	100: {HP: 4000, MP: 10000, Main: 440, Sub: 420, Div: 2780},
}

func (l Level) HP() int {
	return levelMods[l].HP
}

func (l Level) MP() int {
	return levelMods[l].MP
}

func (l Level) Main() int {
	return levelMods[l].Main
}

func (l Level) Sub() int {
	return levelMods[l].Sub
}

func (l Level) Div() int {
	return levelMods[l].Div
}
