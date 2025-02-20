package goprintarray

import(
	"fmt"
)

func PrintArray(arr []int){
	for _, val := range arr{
		fmt.Printf("%d ", val)
	}
}

func PrintMatrix(mat [][]int){
	for _, row := range mat{
		PrintArray(row)
		fmt.Println()
	}
}