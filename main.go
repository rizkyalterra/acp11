package main

import "fmt"

func SimpleEquations(a, b, c int) {
	// your code here
	x, y, z := 1, 1, 1
	// O(c)
	for true {
		if x+y+z > a || x*y*z > b || x*x+y*y+z*z > c {
			fmt.Println("no solution")
			break
		}

		if x+y+z == a {
			if x*y*z == b {
				if x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					break
				} else if z == y {
					z++
				} else if y == x {
					y++
				} else {
					x++
				}
			} else if z == y {
				z++
			} else if y == x {
				y++
			} else {
				x++
			}
		} else if z == y {
			z++
		} else if y == x {
			y++
		} else {
			x++
		}
	}
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(7, 8, 21) // 2, 2, 3
	SimpleEquations(6, 6, 14) // 1 2 3
}
