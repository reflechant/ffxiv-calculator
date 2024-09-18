package main

import "fmt"

func main() {
	gear, _ := runForm()
	bis := gear.BiS(GNB, 100, KeepersOfTheMoon, []*Materia{SavageAim12, SavageMight12, HeavensEye12, QuickArm12}, 2.4, 2.5)

	if bis.Lvl == 0 {
		fmt.Println("BiS with GCD <= 2.5 not found")
		return
	}

	fmt.Printf("BiS with dmg %v:\n", bis.DamageNormalized())
	fmt.Println(bis.String())
}
