package goprintarray

import(
	"fmt"
)

func printArray(arr []int){
	for _, val := range arr{
		fmt.Printf("%d ", val)
	}
}

func printMatrix(mat [][]int){
	for _, row := range mat{
		printArray(row)
		fmt.Println()
	}
}