package main

import (
	"fmt"
)

func main() {
	var R int
	fmt.Scan(&R)

	var L int
	fmt.Scan(&L)

	fmt.Println("------------------------------")

	//create an array of tuples of L lines
	tuples := make([][]int, L)
	//fmt.Println(tuples)
	tuples[0] = append(tuples[0], R)
	//tuples[1] = append(tuples[1], 1, 1)

	var ii int = 0
	for _, i := range tuples {
		//fmt.Println("type of id params = ", reflect.TypeOf(i))
		fmt.Println("slice:", i)
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		//run inside element of the array - tuple

		var count int
		var z int = 0
		var sum int = 0
		for v, j := range i {
			z = j
			//fmt.Println(z)

			fmt.Println("total in slice i: ", len(i))
			fmt.Println("position slice v, j: ", v, j)
			fmt.Println("value en el slice: ", i[v])
			//fmt.Println("value j+1: ", v+1)
			fmt.Println("-------------------")
			//chek not the last one
			if v < len(i) {
				//compare values in the slice
				fmt.Println("v, len[i]: ", v, len(i))
				//fmt.Println("value en el slice: ", tuples[j][v])
				//fmt.Println("value en el slice: ", tuples[j][v+1])
				fmt.Println("value en el slice: ", j)

				//if j == j+1 {
				fmt.Println("the same")
				sum += 1 //sum 1
				tuples[ii+1] = append(tuples[ii+1], sum, z)
				fmt.Println(tuples)

			} else {
				fmt.Println("NOT the same")
				//add (count, number)
				sum = 0
				count = 1 //always is 1
				//number = j //same type
				tuples[ii+1] = append(tuples[ii+1], count, j)
				//i = append(i, count, j)
				//fmt.Println(count, j)
				fmt.Println(tuples)
			}

		}
		ii += 1

		//fmt.Println(sum, z)
		//tuples[ii+1] = append(tuples[ii+1], sum, z)
		//i = append(i, sum, z) //z & sum must be global

	}

	fmt.Println(tuples)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("answer") // Write answer to stdout
}
