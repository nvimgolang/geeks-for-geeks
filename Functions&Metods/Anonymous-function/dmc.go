package main

import "fmt"

// factory function that returns a fighting style for Dante
func DanteStyle(style string) func(enemy string, damage int) string {
	// anonymous function for attacking
	attack := func(enemy string, damage int) string {
		switch style {
		case "Swordmaster":
			return fmt.Sprintf(
				"Dante in Swordmaster style slashes %s for %d damage! 'Stinger'!",
				enemy,
				damage,
			)
		case "Gunslinger":
			return fmt.Sprintf(
				"Dante in Gunslinger style blasts %s for %d damage! 'Bang Bang!'",
				enemy,
				damage,
			)
		default:
			return fmt.Sprintf(
				"Dante casually kicks %s for %d damage, 'Get lost!",
				enemy,
				damage,
			)
		}
	}
	return attack
}

func main() {
	// Dante picks his fighting styles
	swordmaster := DanteStyle(
		"Swordmaster",
	) // Get the Swordmaster style function
	gunslinger := DanteStyle(
		"Gunslinger",
	) // Get the Gunslinger style function

	// Dante wrecks some demons
	fmt.Println(swordmaster("Mundus", 50))  // Attack Mundus
	fmt.Println(gunslinger("Vergil", 30))   // Attack Vergil
	fmt.Println(swordmaster("Phantom", 70)) // Attack Phantom
}
