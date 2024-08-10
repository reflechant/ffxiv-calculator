package main

type Level struct {
	HP   int
	MP   int
	Main int
	Sub  int
	Div  int
}

var (
	Lvl1   = Level{HP: 86, MP: 10000, Main: 20, Sub: 56, Div: 56}
	Lvl100 = Level{HP: 4000, MP: 10000, Main: 440, Sub: 420, Div: 2780}
)
