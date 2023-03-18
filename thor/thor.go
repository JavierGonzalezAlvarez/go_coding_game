package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
	// lightX: the X position of the light of power
	// lightY: the Y position of the light of power
	// initialTx: Thor's starting X position
	// initialTy: Thor's starting Y position
	var lightX, lightY, initialTx, initialTy int
	fmt.Scan(&lightX, &lightY, &initialTx, &initialTy)

	var Tx, Ty int
	var Direction string
	Tx = initialTx
	Ty = initialTy

	for {
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
		var remainingTurns int
		fmt.Scan(&remainingTurns)

		// A single line providing the move to be made: N NE E SE S SW W or NW
		//fmt.Println("SE")

		//fmt.Printf("light position X: %d Y: %d \n", lightX, lightY)
		//fmt.Printf("thor position X: %d Y: %d \n",initialTx, initialTy)

		if (Tx > lightX) && (Ty > lightY) {
			Tx -= 1
			Ty -= -1
			Direction = "NW"
		} else if (Tx > lightX) && (Ty < lightY) {
			Tx -= 1
			Ty += 1
			Direction = "SW"
		} else if (Tx < lightX) && (Ty < lightY) {
			Direction = "SE"
			Tx += 1
			Ty += 1
			Direction = "SE"
		} else if (Tx < lightX) && (Ty > lightY) {
			Direction = "NE"
			Tx += 1
			Ty -= 1
			Direction = "NE"
		} else if Tx > lightX {
			Tx -= 1
			Direction = "W"
		} else if Tx < lightX {
			Tx += 1
			Direction = "E"
		} else if Ty > lightY {
			Ty -= 1
			Direction = "N"
		} else if Ty < lightY {
			Ty += 1
			Direction = "S"
		}

		fmt.Println(Direction)
	}
}
