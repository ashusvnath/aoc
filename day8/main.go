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
	parsedData := Parse(data)
	part1 := Part1(parsedData)
	fmt.Printf("Part1 : %v\n", part1)
}

func Parse(data []byte) [][]int {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	matrix := make([][]int, len(lines))
	for row, lineString := range lines {
		matrix[row] = make([]int, len(lines))
		for col, c := range lineString {
			matrix[row][col] = int(c) - '0'
		}
	}
	return matrix
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
	fmt.Printf("Seen : %v\n", visible)
	return 4*l + len(visible)
}

func Part2(data [][]int) int {
	return 0
}
