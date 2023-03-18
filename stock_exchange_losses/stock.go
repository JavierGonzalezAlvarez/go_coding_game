package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 100000), 100000)
	var inputs []string

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")

	x := []int{}
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 32)
		_ = v

		fmt.Printf(" %d ", v)
		fmt.Println("type v = ", reflect.TypeOf(v))
		//convet int64 to int
		x = append(x, int(v))
	}

	fmt.Println("array", x)
	min := x[0]
	max := x[0]
	loss := x[0]

	i := 0
	j := i + 1
	for _, value := range x {
		fmt.Println("-----------------------------")
		fmt.Println("value: ", value)
		fmt.Println("i init: ", i)
		fmt.Println("j init: ", j)
		if len(x) == j {

		} else {
			if x[i] > x[j] {
				//loss
				loss_local := x[j] - x[i]
				fmt.Println("local_loss: ", loss_local)
				if loss_local < loss {
					loss = loss_local
					fmt.Println("loss: ", loss)
					if value < min {
						min = value
					}
					if value > max {
						max = value
					}
				}
			} else {
				//benefit
				fmt.Println("benefit: ", 0)
				i = j
				j = i
			}
			j = j + 1
			fmt.Println("i: ", i)
			fmt.Println("j: ", j)

		}
	}

	fmt.Println("max value: ", max)
	fmt.Println("min value: ", min)

	var val int = loss
	if loss > 0 {
		val = 0
	} else {
		val = loss
	}
	var answer string = strconv.Itoa(val)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(answer)
}
