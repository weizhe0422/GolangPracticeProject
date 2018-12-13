package temp

import (
	"fmt"
	"os"
)
type point struct {
	i, j int
}

var dirs = [4]point{
	{-1,0},{0,-1},{1,-0},{0,1}}

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

func readMazeFile(filename string) [][]int{
	file, err := os.Open(filename)
	if err!=nil{
		panic(err)
	}

	defer file.Close()

	var col,row int
	fmt.Fscanf(file, "%d %d",&row,&col)

	maze := make([][]int, row)
	for i := range maze{
		maze[i] = make([]int,col)
		for j := range maze[i]{
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}

func walk(maze [][]int, start, end point) ([][]int, int, []point) {
	totalSteps := 0

	steps := make([][]int, len(maze))
	for i := range steps{
		steps[i] = make([]int,len(maze[i]))
	}

	Q := []point{start}
	for len(Q) >0 {
		cur := Q[0]
		Q = Q[1:]

		//Break for loop if arrive END point
		if cur == end{
			break
		}

		for _, dir := range dirs{
			next := cur.add(dir)

			//Remove illegal cases
			//1. Hit Wall
			value, ok := next.at(maze)
			if !ok || value == 1 {
				continue
			}

			//2. Walked already
			value, ok = next.at(steps)
			if !ok || value != 0 {
				continue
			}

			//3. back to start point
			if next == start {
				continue
			}

			//Do things
			//1. Add 1 step  for next step into steps
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			//2. Add next step into Q
			Q = append(Q, next)
			//3. Add totalSteps
			totalSteps = steps[next.i][next.j]
		}
	}

	backIdx := totalSteps
	routeSteps := []point{end}
	curStep := end
	for backIdx > 0 {
		for _, dir := range dirs{
			next:=curStep.add(dir)
			value, ok := next.at(steps)
			if !ok || value != backIdx-1{
				continue
			}
			routeSteps = append(routeSteps,next)
			curStep = next
			backIdx-=1
		}
	}


	return steps,totalSteps,routeSteps
}

func main(){
	maze := readMazeFile("GolangPracticeProject/FromMoocs/Maze/maze.txt")
	for _, row := range maze{
		for _, value := range row{
			fmt.Printf("%3d",value)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("===========================")
	fmt.Println()

	resultMaze,totalSteps,routeSteps := walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})


	for _, row := range resultMaze{
		for _, value := range row{
			fmt.Printf("%3d",value)
		}
		fmt.Println()
	}

	fmt.Printf("Total Steps: %d\n",totalSteps)
	fmt.Println(routeSteps)
}
