package board

import "fmt"

func ExampleMakeBoardRow() {
	row := MakeBoardRow(3, "A")
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	// A
	// A
	// A
}

func ExampleMakeEmptyBoardRow() {
	row := MakeEmptyBoardRow(3)
	for _, v := range row {
		fmt.Println(v)
	}
	// Output:
	//
	//
	//
}
