package board

import "fmt"

func ExampleMakeBoard() {
	board := MakeBoard(3, 2, "A")
	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// [A A A]
	// [A A A]
}

func ExampleMakeEmptyBoard() {
	board := MakeEmptyBoard(3, 2)
	for _, row := range board {
		fmt.Println(row)
	}

	// Output:
	// [     ]
	// [     ]
}

func ExampleGetRow() {
	board := MakeEmptyBoard(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(GetRow(board, 0))
	fmt.Println(GetRow(board, 1))

	// Output:
	// [A B C]
	// [D E F]
}

func ExampleGetColumn() {
	board := MakeEmptyBoard(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(GetColumn(board, 0))
	fmt.Println(GetColumn(board, 1))
	fmt.Println(GetColumn(board, 2))

	// Output:
	// [A D]
	// [B E]
	// [C F]
}

func ExampleGetDiagRight() {
	board := MakeEmptyBoard(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(GetDiagRight(board))

	// Output:
	// [A E I]
}

func ExampleGetDiagLeft() {
	board := MakeEmptyBoard(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(GetDiagLeft(board))

	// Output:
	// [C E G]
}
