package graph

import "container/list"

// public void BFS(Node node) {
// 	Queue data = new ArrayDeque<Node>();
// 	data.add(node);

// 	while(!data.isEmpty()) {
// 			Node curr = (Node) data.poll();
// 			node.marked = true
// 			for(Node child: node.chidren) {
// 					if (!child.marked) {
// 							data.add(child);
// 					}
// 			}
// 	}
// }

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	numIslands := 0
	queue := list.New()

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 49 {
				numIslands++

				grid[r][c] = 48
				queue.PushBack(r*nc + c)

				for queue.Len() > 0 {
					e := queue.Front() // Pop element from queue
					queue.Remove(e)    // Dequeue

					row := e.Value.(int) / nc
					col := e.Value.(int) % nc

					if col-1 >= 0 && grid[row][col-1] == 49 {
						grid[row][col-1] = 48
						queue.PushBack(row*nc + col - 1)
					}

					if col+1 < nc && grid[row][col+1] == 49 {
						grid[row][col+1] = 48
						queue.PushBack(row*nc + col + 1)
					}

					if row-1 >= 0 && grid[row-1][col] == 49 {
						grid[row-1][col] = 48
						queue.PushBack((row-1)*nc + col)
					}

					if row+1 < nr && grid[row+1][col] == 49 {
						grid[row+1][col] = 48
						queue.PushBack((row+1)*nc + col)
					}
				}
			}
		}
	}

	return numIslands
}

func GeneralBFS(grid [][]byte) {
	if grid == nil || len(grid) == 0 {
		return
	}

	nr := len(grid)
	nc := len(grid[0])
	queue := list.New()

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 49 {
				grid[r][c] = 48
				queue.PushBack(r*nc + c)

				for queue.Len() > 0 {
					e := queue.Front() // Pop element from queue
					queue.Remove(e)    // Dequeue

					row := e.Value.(int) / nc
					col := e.Value.(int) % nc

					// 1.) bounds check 2.) search condition
					if col-1 >= 0 && grid[row][col-1] == 49 {
						grid[row][col-1] = 48
						queue.PushBack(row*nc + col - 1)
					}

					if col+1 < nc && grid[row][col+1] == 49 {
						grid[row][col+1] = 48
						queue.PushBack(row*nc + col + 1)
					}

					if row-1 >= 0 && grid[row-1][col] == 49 {
						grid[row-1][col] = 48
						queue.PushBack((row-1)*nc + col)
					}

					if row+1 < nr && grid[row+1][col] == 49 {
						grid[row+1][col] = 48
						queue.PushBack((row+1)*nc + col)
					}
				}
			}
		}
	}

	return
}
