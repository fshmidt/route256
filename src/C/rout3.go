package main

import (
	"fmt"
)

func main() {
	var nId, nMs, b, mem int
	fmt.Scan(&nId, &nMs)
	table := make(map[int]int)
	for a := 1; a <= nMs; a++ {
		var t, id int
		fmt.Scan(&t, &id)
		switch t {
		case 1:
			b += 1
			switch id {
			case 0:
				mem = b
			default:
				table[id] = b
			}
		case 2:
			if table[id] > mem {
				fmt.Println(table[id])
			} else {
				fmt.Println(mem)
			}
		}
	}
}
