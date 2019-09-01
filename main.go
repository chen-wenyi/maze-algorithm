package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	var row, col int
	fmt.Fscanf(file, "%d%d", &row, &col)
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

func (p point) add(k point) point {
	return point{p.i + k.i, p.j + k.j}
}

func (p point) valueAt(place [][]int) (int, bool) {
	if p.i < 0 || p.j < 0 || p.i > len(place)-1 || p.j > len(place[p.i])-1 {
		return 0, false
	}
	return place[p.i][p.j], true
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func makeInitialSteps(maze [][]int) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	return steps
}

func walk(maze [][]int, start, end point) [][]int {
	steps := makeInitialSteps(maze)
	queue := []point{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == end {
			break
		}
		for _, dir := range dirs {
			next := curr.add(dir)
			if val, ok := next.valueAt(maze); !ok {
				continue
			} else {
				// next value at maze == 1
				// next value at steps == start
				if val == 1 || next == start {
					continue
				}
				// next value at steps != 0
				if val, _ := next.valueAt(steps); val != 0 {
					continue
				}
				queue = append(queue, next)
				currValue, _ := curr.valueAt(steps)
				steps[next.i][next.j] = currValue + 1
			}
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze.txt")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
