package main

import "fmt"

func main() {
	fmt.Println(DamageNormalized(Attributes{
		Lvl: LevelMod100,
		Job: JobModifiers["GNB"],
	}, 100))
}
