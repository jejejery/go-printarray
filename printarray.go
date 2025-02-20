package goprintarray

import(
	"fmt"
)

func PrintArray(arr []int){
	fmt.Print("[")
	for _, val := range arr{
		fmt.Printf("%d ", val)
	}
	fmt.Print("]")
}

func PrintStrArray(arr []string){
	fmt.Print("[")
	for _, val := range arr{
		fmt.Printf("%s ", val)
	}
	fmt.Print("]")
}

func PrintMatrix(mat [][]int){
	for _, row := range mat{
		PrintArray(row)
		fmt.Println()
	}
}