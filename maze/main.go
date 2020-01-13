package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) in(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(new [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(new) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(new[p.i]) {
		return 0, false
	}

	return new[p.i][p.j], true
}

//direction
var direct = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//init Q
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		//select redirection
		for _, dir := range direct {
			next := cur.in(dir)

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func showMaze(m [][]int) {
	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func main() {
	maze := readMaze("maze.in")

	//showMaze(maze)

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	showMaze(steps)
}
