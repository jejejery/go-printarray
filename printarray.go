package goprintarray

import(
	"fmt"
)

func PrintArray(arr []int){
	fmt.Print("[")
	for i := 0; i < len(arr); i++{
		fmt.Printf("%d", arr[i])
		if i != len(arr)-1{
			fmt.Print(" ")
		}
	}
	fmt.Print("]")
}

func PrintStrArray(arr []string){
	fmt.Print("[")
	for i := 0; i < len(arr); i++{
		fmt.Printf("%s", arr[i])
		if i != len(arr)-1{
			fmt.Print(" ")
		}
	}
	fmt.Print("]")
}

func PrintMatrix(mat [][]int){
	for _, row := range mat{
		PrintArray(row)
		fmt.Println()
	}
}