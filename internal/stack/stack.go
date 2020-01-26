package stack

import (
	"container/list"
	"fmt"
)

// https://golang.org/pkg/container/list/#List.PushFront

func Stack() {
	stack := list.New()
	stack.PushFront(" world")
	stack.PushFront("hello")

	for stack.Len() > 0 {
		e := stack.Front() // First element
		fmt.Print(e.Value)

		stack.Remove(e) // Dequeue
	}
}

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	numIslands := 0
	stack := list.New()

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 49 {
				numIslands++

				grid[r][c] = 48
				stack.PushFront(r*nc + c)

				for stack.Len() > 0 {
					e := stack.Front() // Pop element from queue
					stack.Remove(e)    // Dequeue

					row := e.Value.(int) / nc
					col := e.Value.(int) % nc

					if col-1 >= 0 && grid[row][col-1] == 49 {
						grid[row][col-1] = 48
						stack.PushFront(row*nc + col - 1)
					}

					if col+1 < nc && grid[row][col+1] == 49 {
						grid[row][col+1] = 48
						stack.PushFront(row*nc + col + 1)
					}

					if row-1 >= 0 && grid[row-1][col] == 49 {
						grid[row-1][col] = 48
						stack.PushFront((row-1)*nc + col)
					}

					if row+1 < nr && grid[row+1][col] == 49 {
						grid[row+1][col] = 48
						stack.PushFront((row+1)*nc + col)
					}
				}
			}
		}
	}

	return numIslands
}

// do a 2D traversal through the array in the outer loop
// once you find an island mark it as 0 and then kick off the inner loop
// inner loop will queue up all neighbors of current index and if it is land then it will queue it up
// inner loop will mark all land it iterates through as 0. inner loop will terinate once queue is empty.
// once inner loop is terminated resume outer loop. continue up 2D traversal is complete.
