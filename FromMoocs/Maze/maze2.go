package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add (r point) point{
	return point{p.i+r.i,p.j+r.j}
}

func (p point) at(grid [][]int) (int, bool){
	if p.i < 0 || p.i >= len(grid){
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]){
		return 0, false
	}

	return grid[p.i][p.j], true
}

var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1}}

func readfile(filename string) [][]int{
	file, err := os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	var row,col int
	fmt.Fscanf(file,"%d %d",&row,&col)
	maze := make([][]int,row)
	for i := range maze{
		maze[i] = make([]int,col)
		for j := range maze[i]{
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}


func walk(maze [][]int,start, end point) ([][]int, int){
	steps := make([][]int, len(maze))
	totalsteps := 0

	for i := range steps{
		steps[i] = make([]int,len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end{
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			value, ok := next.at(maze)
			if !ok || value == 1 {
				continue
			}

			value, ok = next.at(steps)
			if !ok || value != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			Q = append(Q,next)
			totalsteps = steps[next.i][next.j]
		}
	}
	return steps,totalsteps
}

func main() {
	maze := readfile("GolangPracticeProject/FromMoocs/Maze/maze.txt")
	for _, row := range maze {
		for _, value := range row{
			fmt.Printf("%3d",value)
		}
		fmt.Println()
	}

	fmt.Println("--------------------------")

	result, totalsteps := walk(maze,point{0,0},point{len(maze),len(maze[0])})
	for _, row := range result {
		for _, value := range row{
			fmt.Printf("%3d",value)
		}
		fmt.Println()
	}

	fmt.Print(totalsteps)
}
