package main

import (
	"fmt"
	"os"
)

func PrintMaze(maze [][]int) {
	for i, _ := range maze {
		for j, _ := range maze[i]{
			fmt.Printf("%3d ", maze[i][j])
		}
		fmt.Println()
	}
}

func ReadMaze(filename string)  [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int

	n, err := fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	maze_info := make([][]int, row)

	for i := range maze_info {
		maze_info[i] = make([]int, col)
		for j := range maze_info[i]{
			fmt.Fscanf(file, "%d", &maze_info[i][j])
		}
	}

	//PrintMaze(maze_info)

	return maze_info


}

type point struct {
	i, j int
}

var move_step = [4]point{
	{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func (p point) add(m point) point {
	return point{(p.i+m.i), (p.j+m.j)}
}

func (p point) at(maze [][]int) (int, bool){
	if p.i < 0 || p.i > len(maze)-1 {
		return 0, false
	}

	if p.j < 0 || p.j > len(maze[0])-1 {
		return 0, false
	}

	return maze[p.i][p.j], true
}

func Walk_maze(maze [][]int, start, end point) ([][]int, int) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	q := []point{start}
	var cur_var int

	for len(q) > 0 {
		cur_step := q[0]
		q = q[1:]

		for _, move := range move_step {
			next := cur_step.add(move)

			val, ok := next.at(maze)
			if !ok || val == 1{
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0{
				continue
			}

			if next == start {
				continue
			}

			q = append(q, next)



			val, ok1 := cur_step.at(steps)
			if ok1 {
				steps[next.i][next.j] = val + 1
				cur_var = val + 1
			}

			q = append(q, next)

		}

	}

	return steps, cur_var

}

func GetWalk(steps [][]int, end point)  [][]int {

	var step_list []point
	cur_point := end
	cur_val := steps[end.i][end.j]

	for cur_val > 0 {
		step_list = append(step_list, cur_point)
		for _, move := range move_step {
			next := cur_point.add(move)

			val, ok := next.at(steps)
			if ok && val == cur_val-1 {
				cur_val = val
				cur_point = next
				break
			}

		}
	}

	step_map := make([][]int, len(steps))
	for i := range step_map {
		step_map[i] = make([]int, len(steps[i]))
	}

	for _, v := range step_list {
		i,j := v.i,v.j
		step_map[i][j] = 1
	}

	return step_map

}

func main() {
	source_maze := ReadMaze("./maze/maze.txt")
	start_point := point{0,0}
	end_point := point{len(source_maze)-1, len(source_maze[0])-1}
	steps, steps_len := Walk_maze(source_maze, start_point, end_point)

	//PrintMaze(steps)
	fmt.Println(steps_len)


	step_map := GetWalk(steps, end_point)
	PrintMaze(step_map)


	//for i := range steps {
	//	for _,val := range steps[i] {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}

}
