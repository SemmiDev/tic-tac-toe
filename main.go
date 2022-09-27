package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	clearScreen()
	fmt.Println("⇝ Mari bermain Silang-Bulat-Silang!")
	fmt.Println("⇝ Anda dapat memilih kotak 1-9 untuk mulai bermain")
	fmt.Println("⇝ Anda akan bermain dengan komputer")
	fmt.Println("⇝ Komputer simbolnya adalah O")
	fmt.Println("⇝ Anda simbolnya adalah X")
	fmt.Println()

	var templateBoard, board [9]string
	for i := 0; i < 9; i++ {
		board[i] = " "
		templateBoard[i] = fmt.Sprintf("%d", i+1)
	}

	printBoard(templateBoard)
	fmt.Println()
	fmt.Println("⇝ Permainan dimulai dalam 5 detik")
	time.Sleep(3 * time.Second)
	clearScreen()

	printBoard(board)
	fmt.Println()

	for {
		clearScreen()
		printBoard(board)

		fmt.Println()
		fmt.Print("Pilih kotak: ")
		var input int
		fmt.Scanln(&input)

		if input < 1 || input > 9 {
			fmt.Println("g boleh. input harus berupa angka dan harus direntang 1-9")
			time.Sleep(time.Second)
			continue
		}

		if !mustKotakEmptyIn(board, input-1) {
			fmt.Println("g boleh. kotak ", input, "udah diisi")
			time.Sleep(time.Second)
			continue
		}

		board[input-1] = "X"

		if isKotakFull(board) {
			theWinner := decideWinner(theWinner(board))
			clearScreen()
			printBoard(board)

			if theWinner != "" {
				fmt.Println(theWinner)
			} else {
				fmt.Println("SERI! (TIDAK ADA PEMENANG)")
			}
			break
		}

		for {
			r := randomInt(1, 9) - 1
			if mustKotakEmptyIn(board, r) {
				board[r] = "O"
				break
			}
		}

		theWinner := decideWinner(theWinner(board))
		if theWinner != "" {
			clearScreen()
			printBoard(board)
			fmt.Println(theWinner)
			break
		}
	}
}

// init function will be called before main function
func init() {
	// set random seed for randomInt function, it usefully for generate random number
	rand.Seed(time.Now().Unix())
}

// randomInt function will return random number in range min and max
func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// clearScreen function will clear screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// printBoard function will print board
func printBoard(board [9]string) {
	fmt.Println(" ", board[0], " | ", board[1], " | ", board[2])
	fmt.Println("_____|_____|_____")
	fmt.Println(" ", board[3], " | ", board[4], " | ", board[5])
	fmt.Println("_____|_____|_____")
	fmt.Println(" ", board[6], " | ", board[7], " | ", board[8])
	fmt.Println("     |     |")
}

// mustKotakEmptyIn function will check if the board is empty in index
func mustKotakEmptyIn(board [9]string, index int) bool {
	return board[index] == " "
}

// isKotakFull function will check if the board is full
func isKotakFull(board [9]string) bool {
	for _, v := range board {
		if v == " " {
			return false
		}
	}
	return true
}

// decideWinner function will decide the winner based on the winner
func decideWinner(winner string) string {
	switch winner {
	case "X":
		return "Anda menang!"
	case "O":
		return "Komputer menang!"
	default:
		return ""
	}
}

// theWinner function will return the winner
func theWinner(board [9]string) (winner string) {
	// check horizontal
	for i := 0; i < 9; i += 3 {
		if board[i] == board[i+1] && board[i+1] == board[i+2] && board[i] != " " {
			return board[i]
		}
	}

	// check vertical
	for i := 0; i < 3; i++ {
		if board[i] == board[i+3] && board[i+3] == board[i+6] && board[i] != " " {
			return board[i]
		}
	}

	// check diagonal
	if board[0] == board[4] && board[4] == board[8] && board[0] != " " {
		return board[0]
	}

	if board[2] == board[4] && board[4] == board[6] && board[2] != " " {
		return board[2]
	}

	return ""
}
