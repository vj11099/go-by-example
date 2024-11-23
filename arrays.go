package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty:", a)
	fmt.Println()

	a[4] = 100
	fmt.Println("new arr:", a)
	fmt.Println("arr by index 4:", a[4])
	fmt.Println()

	fmt.Println("Length of the array above:", len(a))
	fmt.Println()

	b := [5]int{12, 6, 3, 19, 8}
	fmt.Println("declared array:", b)
	fmt.Println()

	b = [...]int{12, 6, 3, 19, 8}
	fmt.Println("auto adjust array length using [...]:", b)
	fmt.Println()

	b = [...]int{12, 4: 9}
	fmt.Println("zeroing till [index]-th value using ':' :", b)
	fmt.Println()

	twoD := [2][3]int{{2, 45, 16}, {12, 37, 62}}
	fmt.Println("2-D Array: ", twoD)
	fmt.Println()

	var new2D [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			new2D[i][j] = fmt.Sprintf("%d%d", i, j)
		}
	}
	fmt.Println("2-D Array: ", new2D)
	fmt.Println()

}
