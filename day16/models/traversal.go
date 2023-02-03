package models

type Traversal struct {
	opened      []string
	current     string
	released    int
	currentRate int
	steps       int
	visited     []string
}

func (t *Traversal) Current() string {
	return t.current
}

func (t *Traversal) Steps() int {
	return t.steps
}

func (t *Traversal) TotalReleased() int {
	return t.released
}

func (t *Traversal) AllValvesOpened() bool {
	return len(t.opened) == RoomsWithNonZeroReleaseRate
}

func (t *Traversal) CompleteSteps(n int) {
	t.released += t.currentRate * (n - t.steps)
	t.steps = n
}

func (t *Traversal) TooManyUnopened() bool {
	return (len(t.visited)-len(t.opened))/len(AllRooms) >= 1
}

func (t *Traversal) VisitAndOpen(toVisit string) *Traversal {
	if contains(t.opened, toVisit) || AllRooms[toVisit].ReleaseRate() == 0 {
		//log.Printf("Step %d, Valve in Room %s already opened. Skipping VisitAndOpen.", t.steps, toVisit)
		return nil
	}
	//log.Printf("Step %d: Visiting room %v and Opening valve", t.steps, toVisit)

	newVisited := append([]string{}, t.opened...)
	newVisited = append(newVisited, toVisit)
	return &Traversal{
		opened:      newVisited,
		current:     toVisit,
		released:    t.released + t.currentRate + t.currentRate,
		currentRate: t.currentRate + AllRooms[toVisit].ReleaseRate(),
		steps:       t.steps + 2,
		visited:     append(t.visited, toVisit),
	}
}

func (t *Traversal) Visit(toVisit string) *Traversal {
	//log.Printf("Step %d: Visiting room %v", t.steps, toVisit)

	return &Traversal{
		opened:      append([]string{}, t.opened...),
		current:     toVisit,
		released:    t.released + t.currentRate,
		currentRate: t.currentRate,
		steps:       t.steps + 1,
		visited:     append(t.visited, toVisit),
	}
}

func NewTraversal(start string) *Traversal {
	return &Traversal{
		opened:      []string{},
		current:     start,
		released:    0,
		currentRate: 0,
		steps:       2,
		visited:     []string{start},
	}
}

func contains(collection []string, elem string) bool {
	for _, e := range collection {
		if e == elem {
			return true
		}
	}
	return false
}
