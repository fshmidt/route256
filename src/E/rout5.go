package main

import (
	"fmt"
)

func main() {
	var nom int
	fmt.Scan(&nom)
	for a := 0; a < nom; a++ {
		var d, prev, no int
		fmt.Scan(&d)
		m := make(map[int]int)
		for i := 0; i < d; i++ {
			if d < 3 || d > 50000 {
				no++
				break
			}
			var ind int
			fmt.Scan(&ind)
			if ind < 1 {
				no++
				break
			}
			if m[ind] == 0 || prev == ind {
				m[ind]++
				prev = ind
			} else {
				no++
				break
			}
		}
		if no > 0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
		if d == 0 {
			fmt.Scan(&d)
		}
	}

}
