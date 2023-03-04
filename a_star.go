package a_star

import (
	"fmt"
	"math"
	"sort"
)

const (
	road = 0
	wall = 1

	posA = -1
	posB = -2

	path = -3
)

const (
	up = iota
	down
	left
	right
)

var direction = map[int]Pos{
	up:    {+0, +1},
	down:  {+0, -1},
	left:  {-1, +0},
	right: {+1, +0},
}

type Pos struct {
	X int
	Y int
}

func AStar(start Pos, end Pos, matrix [][]int) (isAccessible bool, search [][]int) {
	matrix[start.Y][start.X] = posA
	matrix[end.Y][end.X] = posB

	search = make([][]int, len(matrix))
	for y := range search {
		search[y] = make([]int, len(matrix[0]))
		for x := range search[y] {
			search[y][x] = 99
		}
	}

	search[start.Y][start.X] = 0

	var liveList []Pos
	var diedList []Pos
	liveList = append(liveList, start)

	count := 0
	// 开始搜索
	for ; ; count++ {

		//// 查看
		//PrintSearch(search)

		// 判断是否到达目标点
		if search[end.Y][end.X] != 99 {
			break
		}

		// 判断存在选项
		if len(liveList) == 0 {
			break
		}

		// 排序
		sort.Slice(liveList, func(i, j int) bool {
			return math.Abs(float64(end.X-liveList[i].X))+math.Abs(float64(end.Y-liveList[i].Y)) <
				math.Abs(float64(end.X-liveList[j].X))+math.Abs(float64(end.Y-liveList[j].Y))
		})

		livePos := liveList[0]

		// 周围查找
		for i := up; i <= right; i++ {
			dirc := direction[i]
			step := Pos{X: livePos.X + dirc.X, Y: livePos.Y + dirc.Y}

			// 判断可达
			if matrix[step.Y][step.X] == wall {
				continue
			}

			// 更新到达该点的最小消耗
			if search[step.Y][step.X] > search[livePos.Y][livePos.X]+1 {
				search[step.Y][step.X] = search[livePos.Y][livePos.X] + 1
				// 更新时判断是否已被加入候选
				if isInList(step, liveList) || isInList(step, diedList) {
					continue
				}
				// 加入候选
				liveList = append(liveList, step)
			}
		}

		// 搜索完毕后改为死亡点
		liveList = liveList[1:]
		diedList = append(diedList, livePos)
	}

	isAccessible = true
	//PrintMatrix(matrix)
	//PrintSearch(search)
	//PrintStep(end, matrix, search)
	//fmt.Println("查找了 ", count, " 次")

	return
}

func isInList(p Pos, list []Pos) (isIn bool) {
	for _, p2 := range list {
		if p == p2 {
			return true
		}
	}
	return
}

func PrintMatrix(matrix [][]int) {
	print("↓ →")
	for i := 0; i < len(matrix[0]); i++ {
		fmt.Printf("%3d", i)
	}
	println()
	for y, ints := range matrix {
		fmt.Printf("%3d", y)
		for _, code := range ints {
			switch code {
			case road:
				print("   ")
			case wall:
				print("███")
			case posA:
				print("→⊠←")
			case posB:
				print("→⊡←")
			case path:
				print(" O ")
			}
		}
		println()
	}

	for i := 0; i <= len(matrix[0]); i++ {
		print("+++")
	}
	println()
}

func PrintSearch(search [][]int) {
	print("↓ →")
	for i := 0; i < len(search[0]); i++ {
		fmt.Printf("%3d", i)
	}
	println()
	for y, ints := range search {
		fmt.Printf("%3d", y)
		for _, code := range ints {
			switch code {
			case 99:
				print("---")
			default:
				fmt.Printf("%3d", code)
			}
		}
		println()
	}

	for i := 0; i <= len(search[0]); i++ {
		print("+++")
	}
	println()
}

func PrintStep(end Pos, matrix, search [][]int) {
	value := search[end.Y][end.X]
	matrix[end.Y][end.X] = path

PATH:
	for {
		if value == 0 {
			break
		}
		for dire := up; dire <= right; dire++ {
			if search[end.Y+direction[dire].Y][end.X+direction[dire].X] == value-1 {
				value -= 1
				end.Y += direction[dire].Y
				end.X += direction[dire].X
				matrix[end.Y][end.X] = path
				continue PATH
			}
		}
		panic("error")
	}

	PrintMatrix(matrix)
}
