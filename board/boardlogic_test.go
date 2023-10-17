package board

import "fmt"

func ExampleBoardRowContainsOnly() {
	board := MakeEmptyBoard(3, 2)
	board[0][0] = "A"
	board[0][1] = "A"
	board[0][2] = "A"
	board[1][0] = "D"
	board[1][1] = "E"
	board[1][2] = "F"

	fmt.Println(BoardRowContainsOnly(board, 0, "A"))
	fmt.Println(BoardRowContainsOnly(board, 0, "B"))
	fmt.Println(BoardRowContainsOnly(board, 1, "D"))
	fmt.Println(BoardRowContainsOnly(board, 1, "E"))

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleBoardColumnContainsOnly() {
	board := MakeEmptyBoard(3, 2)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "A"
	board[1][1] = "E"
	board[1][2] = "C"

	fmt.Println(BoardColumnContainsOnly(board, 0, "A"))
	fmt.Println(BoardColumnContainsOnly(board, 0, "B"))
	fmt.Println(BoardColumnContainsOnly(board, 1, "B"))
	fmt.Println(BoardColumnContainsOnly(board, 1, "E"))
	fmt.Println(BoardColumnContainsOnly(board, 2, "C"))
	fmt.Println(BoardColumnContainsOnly(board, 2, "F"))

	// Output:
	// true
	// false
	// false
	// false
	// true
	// false
}

func ExampleBoardDiagRightContainsOnly() {
	board := MakeEmptyBoard(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "A"
	board[1][2] = "F"
	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "A"

	fmt.Println(BoardDiagRightContainsOnly(board, "A"))
	fmt.Println(BoardDiagRightContainsOnly(board, "B"))

	// Output:
	// true
	// false
}

func ExampleBoardDiagLeftContainsOnly() {
	board := MakeEmptyBoard(3, 3)
	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"
	board[1][0] = "D"
	board[1][1] = "C"
	board[1][2] = "F"
	board[2][0] = "C"
	board[2][1] = "H"
	board[2][2] = "I"

	fmt.Println(BoardDiagLeftContainsOnly(board, "C"))
	fmt.Println(BoardDiagLeftContainsOnly(board, "B"))

	// Output:
	// true
	// false
}

func ExampleBoardContainsOnly() {
	board1 := MakeEmptyBoard(3, 2)
	fmt.Println(BoardContainsOnly(board1, " "))
	fmt.Println(BoardContainsOnly(board1, "A"))
	board1[0][0] = "A"
	fmt.Println(BoardContainsOnly(board1, " "))

	board2 := MakeBoard(3, 2, "A")
	fmt.Println(BoardContainsOnly(board2, "A"))
	fmt.Println(BoardContainsOnly(board2, "B"))
	board2[0][0] = "B"
	fmt.Println(BoardContainsOnly(board2, "A"))

	// Output:
	// true
	// false
	// false
	// true
	// false
	// false
}

func ExampleBoardDoesNotContain() {
	board1 := MakeEmptyBoard(3, 2)
	fmt.Println(BoardDoesNotContain(board1, " "))
	fmt.Println(BoardDoesNotContain(board1, "A"))
	board1[0][0] = "A"
	fmt.Println(BoardDoesNotContain(board1, " "))
	fmt.Println(BoardDoesNotContain(board1, "A"))
	board1[0][1] = "B"
	board1[0][2] = "C"
	board1[1][0] = "D"
	board1[1][1] = "E"
	board1[1][2] = "F"
	fmt.Println(BoardDoesNotContain(board1, " "))
	fmt.Println(BoardDoesNotContain(board1, "A"))
	fmt.Println(BoardDoesNotContain(board1, "B"))
	fmt.Println(BoardDoesNotContain(board1, "K"))

	// Output:
	// false
	// true
	// false
	// false
	// true
	// false
	// false
	// true
}
