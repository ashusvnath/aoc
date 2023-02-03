package main

import (
	"day16/models"
	"day16/parser"
	"day16/utility"
	"flag"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var filepath string
var verbose bool
var cpuprofile string

func init() {
	flag.BoolVar(&verbose, "v", false, "show debug logs")
	flag.StringVar(&filepath, "f", "test.txt", "path to file")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
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
	flag.Parse()
	if !verbose {
		log.SetOutput(io.Discard)
	}

	data := readFile(filepath)
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	log.Printf("Data:\n%v", string(data))
	parser.Parse(strings.TrimRight(string(data), "\n"))
	log.Printf("AllRooms:\n %v", models.AllRooms)
	log.Printf("Part1: %d\n", Part1())
}

func Part1() int {
	traversal := models.NewTraversal("AA")
	q := utility.NewQueue[*models.Traversal]()
	q.Enqueue(traversal)
	maxSteps := 30
	completed := []*models.Traversal{}
	seenSteps := -1
	iter := 1
	for ; q.Len() > 0; iter++ {
		current := q.Dequeue()
		if current.Steps() > seenSteps {
			seenSteps = current.Steps()
			log.Printf("Traversal with %d steps found at iteration %d", seenSteps, iter)
		}
		connectedRooms := models.AllRooms[current.Current()].ConnectedRoomIds()
		if current.AllValvesOpened() {
			current.CompleteSteps(maxSteps)
			log.Printf("Found completed: %v", current)
			completed = append(completed, current)
			continue
		}
		// if current.TooManyUnopened() {
		// 	log.Printf("More rooms visited than opened")
		// 	continue
		// }
		atLeastOneEnqueued := false
		for _, roomId := range connectedRooms {
			traversals := []*models.Traversal{current.Visit(roomId)}
			if models.AllRooms[roomId].ReleaseRate() > 0 {
				traversals = []*models.Traversal{traversals[0], current.VisitAndOpen(roomId)}
			}
			for _, next := range traversals {
				if next == nil {
					continue
				}
				if next.Steps() >= maxSteps {
					completed = append(completed, next)
				} else {
					atLeastOneEnqueued = true
					q.Enqueue(next)
				}
			}
		}
		if !atLeastOneEnqueued {
			log.Printf("No new element enqueued at %v.", current)
		}
	}
	max := -1
	var selectedTraversal *models.Traversal
	for _, t := range completed {
		if t.TotalReleased() > max {
			selectedTraversal = t
			max = t.TotalReleased()
		}
	}
	log.Printf("Selected traversal :%v", *selectedTraversal)
	return max
}
