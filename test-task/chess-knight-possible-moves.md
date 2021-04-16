# Chess knight possible moves


```
package main 

import (
	"fmt"
)

// size
const N = 8
const M = 8

var letterMap = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

func countPossibleMoves(mat [N][M]int, p, q int) int {
	var Xmoves = []int{ 2, 1, -1, -2, -2, -1, 1, 2 }
	var Ymoves = []int{ 1, 2, 2, 1, -1, -2, -2, -1 }
	var count = 0

	for i := 0; i < N; i++ {
		var x = p + Xmoves[i];
		var y = q + Ymoves[i]; 

		if x >= 0 && y >= 0 && x < N && y<M && mat[x][y] == 0 {
			count++;
			fmt.Println(x, y) // possible move
		}
	}

	return count
}


func main() {


	var board = [N][M]int{}

	var p = 0
	var q = letterMap["A"]

	fmt.Println("Moves: ", countPossibleMoves(board, p, q))
}
```
