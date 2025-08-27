package main

import (
	"errors"
	"fmt"
)

const (
	Empty   rune = 0
	PlayerO rune = 'O'
	PlayerX rune = 'X'
)

func CheckWinner(b [3][3]rune) (bool, rune) {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, // rows
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, // cols
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, // diags
		{{0, 2}, {1, 1}, {2, 0}},
	}
	for _, ln := range lines {
		a := b[ln[0][0]][ln[0][1]]
		if a != Empty &&
			a == b[ln[1][0]][ln[1][1]] &&
			a == b[ln[2][0]][ln[2][1]] {
			return true, a
		}
	}
	return false, 0
}

func printBoard(b [3][3]rune) {
	fmt.Println()
	fmt.Println("    1   2   3")
	rows := []string{"1", "2", "3"}
	for r := 0; r < 3; r++ {
		a := disp(b[r][0])
		c := disp(b[r][1])
		d := disp(b[r][2])
		fmt.Printf("%s   %c | %c | %c\n", rows[r], a, c, d)
		if r < 2 {
			fmt.Println("   ---+---+---")
		}
	}
	fmt.Println()
}

func disp(v rune) rune {
	if v == 0 {
		return ' '
	}
	return v
}

func TicTacToe() {
	var player1, player2 string
	fmt.Println("Player 1 name: ")
	fmt.Scan(&player1)
	fmt.Println("Player 2 name: ")
	fmt.Scan(&player2)
	// fmt.Printf("%s %s\n", player1, player2)
	fmt.Printf("Player 1 %s is assigned O\nPlayer 2 %s is assigned X\n", player1, player2)
	var ans [3][3]rune
	total_count := 0

	for total_count < 9 {
		var curr_player string
		var inp_char rune
		if total_count%2 == 0 {
			curr_player = player1
			inp_char = 'O'
		} else {
			curr_player = player2
			inp_char = 'X'
		}
		var temp_i, temp_j int
		fmt.Printf("%s's turn: \n", curr_player)
		fmt.Println("Enter row nummber: ")
		fmt.Scan(&temp_i)
		fmt.Println("Enter column nummber: ")
		fmt.Scan(&temp_j)
		if temp_i < 3 && temp_j < 3 {
			ans[temp_i][temp_j] = inp_char
		} else {
			_ = errors.New("index out of range")
		}

		if won, _ := CheckWinner(ans); won {
			printBoard(ans)
			fmt.Printf("Hurray!!! %s won", curr_player)
			break
		}
		total_count++
	}
}
