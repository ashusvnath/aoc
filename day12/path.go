package main

type Path struct {
	path       []complex128
	len        int
	currentIdx complex128
	start      complex128
}

func NewPath(g *Grid, start complex128) *Path {
	return &Path{
		path:       []complex128{start},
		len:        0,
		currentIdx: start,
		start:      start,
	}
}

func (p *Path) Len() int {
	return p.len
}

func (p *Path) Start() complex128 {
	return p.start
}

func (p *Path) Add(idx complex128) *Path {
	newPath := append([]complex128{}, p.path...)
	newPath = append(newPath, idx)
	return &Path{
		path:       newPath,
		len:        p.len + 1,
		currentIdx: idx,
		start:      p.start,
	}
}
