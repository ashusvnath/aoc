package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

type Set[T comparable] map[T]bool

func (s *Set[T]) Contains(val T) bool {
	_, ok := (*s)[val]
	return ok
}

func (s *Set[T]) Add(elem T) {
	(*s)[elem] = true
}

func (s *Set[T]) AsSlice() []T {
	result := make([]T, len(*s))
	idx := 0
	for k := range *s {
		result[idx] = k
		idx++
	}
	return result
}

var filepath string
var verbose bool
var cpuprofile string

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "test.txt", "path to file")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if !verbose {
		log.SetOutput(io.Discard)
	}

}

func readFile(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Could not read contents of %s : %v", filepath, err)
		os.Exit(1)
	}
	return data
}

func main() {
	data := readFile(filepath)
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	mat, idxsByHeight := Parse(data)
	part1 := Part1(mat)
	part2 := Part2(mat, idxsByHeight)
	fmt.Printf("Part1 : %v\nPart2 : %v\n", part1, part2)
}

func Parse(data []byte) ([][]int, map[int][]int) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	l := len(lines)
	matrix := make([][]int, l)
	idxsByHeight := make(map[int][]int)
	for row, lineString := range lines {
		matrix[row] = make([]int, l)
		for col, c := range lineString {
			height := int(c) - '0'
			matrix[row][col] = height
			idx := row*l + col
			if row == 0 || col == 0 || row == l-1 || col == l-1 {
				continue
			}
			idxsByHeight[height] = append(idxsByHeight[height], idx)
		}
	}
	return matrix, idxsByHeight
}

func Part1(data [][]int) int {
	l := len(data) - 1
	visible := make(Set[int])

	for i := 1; i < l; i++ {
		lmax := data[i][0]
		rmax := data[i][l]
		tmax := data[0][i]
		bmax := data[l][i]
		for j := 1; j < l; j++ {
			if data[i][j] > lmax {
				lmax = data[i][j]
				visible.Add(i*(l+1) + j)
			}

			if data[i][l-j] > rmax {
				rmax = data[i][l-j]
				visible.Add(i*(l+1) + l - j)
			}

			if data[j][i] > tmax {
				tmax = data[j][i]
				visible.Add(j*(l+1) + i)
			}

			if data[l-j][i] > bmax {
				bmax = data[l-j][i]
				visible.Add((l-j)*(l+1) + i)
			}
		}
	}
	return 4*l + len(visible)
}

func Part2(data [][]int, idxsByHeight map[int][]int) int {
	maxScenicScore := -1
	maxIdx := -1
	for height := 9; height >= 1; height-- {
		idxs := idxsByHeight[height]
		if idxs == nil {
			continue
		}
		log.Printf("Checking height: %v, Idxs: %v", height, idxs)
		for _, idx := range idxs {
			score, parts := scenicScore(idx, data)
			if score > maxScenicScore {
				log.Printf("Part2: Max score beaten at: Idx: %d, score: %d, parts: %v", idx, score, parts)
				maxScenicScore = score
				maxIdx = idx
			}
		}
	}
	log.Printf("Max idx: %v", maxIdx)
	return maxScenicScore
}

func scenicScore(idx int, mat [][]int) (int, [4]int) {
	l := len(mat)
	x, y := idx/l, idx%l
	h := mat[x][y]
	checkRight, checkUp, checkLeft, checkDown := true, true, true, true
	rightScore, upScore, leftScore, downScore := l-y-1, x, y, l-x-1
	for i := 1; i < l && (checkLeft || checkDown || checkRight || checkUp); i++ {
		if checkDown && x+i < l && mat[x+i][y] >= h {
			downScore = i
			checkDown = false
		}
		if checkRight && y+i < l && mat[x][y+i] >= h {
			rightScore = i
			checkRight = false
		}
		if checkUp && x-i > -1 && mat[x-i][y] >= h {
			upScore = i
			checkUp = false
		}
		if checkLeft && y-i > -1 && mat[x][y-i] >= h {
			leftScore = i
			checkLeft = false
		}
	}
	score := leftScore * rightScore * upScore * downScore
	return score, [4]int{upScore, leftScore, rightScore, downScore}
}
