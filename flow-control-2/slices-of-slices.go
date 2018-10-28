
// Slices of slices

// Slices can contain any type, including other slices.
package main

import "fmt"
import "strings"

func main(){
	//create a tic-tac board.
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	//The Players take turns.
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	for i := 0; i<len(board); i++ {
		fmt.Printf("%s/n", strings.Join(board[i]," "))
	}
}

	