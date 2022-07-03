package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nom int
	fmt.Fscan(in, &nom)
	for a := nom; a > 0; a-- {
		var x int
		fmt.Fscan(in, &x)
		arr := make([][2]int, x)
		for i := range arr {
			fmt.Fscan(in, &arr[i][0])
			arr[i][1] = i + 1
		}

		for i := x / 2; i > 0; i-- {
			fmt.Printf("%d ", arr[0][1])
			var minDif, index int
			for j := 1; j < len(arr); j++ {
				if j == 1 || (arr[j][0]-arr[0][0])*(arr[j][0]-arr[0][0]) < minDif*minDif {
					minDif = arr[j][0] - arr[0][0]
					index = j
				}
			}
			fmt.Println(arr[index][1])
			newArr := make([][2]int, len(arr)-2)
			if index != 1 && index != len(arr)-1 {
				newArr = append(arr[1:index], arr[index+1:]...)
			} else if index != len(arr)-1 {
				newArr = arr[index+1:]
			} else {
				newArr = arr[1:index]
			}
			arr = newArr
		}
		fmt.Printf("\n")
	}
}
