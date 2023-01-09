package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var cdCommandPatternString string = `^\$ cd (?P<name>[a-zA-Z/.]+)$`
var lsCommandPatternString string = `^\$ ls$`
var dirOutputLinePatternString string = `^dir (?P<name>[a-zA-Z.]+)$`
var fileOutputLinePatternString string = `^(?P<size>\d+) (?P<name>[a-zA-Z.]+)`
var cdCommandPattern *regexp.Regexp
var lsCommandPattern *regexp.Regexp
var dirOutputLinePattern *regexp.Regexp
var fileOutputLinePattern *regexp.Regexp
var knownDirs []*AOCDir

func init() {
	cdCommandPattern = regexp.MustCompile(cdCommandPatternString)
	lsCommandPattern = regexp.MustCompile(lsCommandPatternString)
	dirOutputLinePattern = regexp.MustCompile(dirOutputLinePatternString)
	fileOutputLinePattern = regexp.MustCompile(fileOutputLinePatternString)
	knownDirs = make([]*AOCDir, 0)
}

type Entry interface {
	Name() string
	Size() int
}

type AOCDir struct {
	name    string
	entries []Entry
	size    int
	parent  *AOCDir
}

func (d *AOCDir) Name() string {
	return d.name
}

func (d *AOCDir) Size() int {
	return d.size
}

func (d AOCDir) String() string {
	return fmt.Sprintf("- %s (dir, size=%d)", d.name, d.size)
}

func NewDir(name string, parent *AOCDir) *AOCDir {
	dir := &AOCDir{name, []Entry{}, 0, parent}
	if parent == nil {
		dir.parent = dir
	}
	knownDirs = append(knownDirs, dir)
	return dir
}

func (d *AOCDir) AddEntry(entry Entry) {
	d.entries = append(d.entries, entry)
	d.updateSize(entry.Size())
}

func (d *AOCDir) updateSize(s int) {
	d.size += s
	if d.parent != nil && d.parent != d {
		d.parent.updateSize(s)
	}
}

func (d *AOCDir) find(name string) *AOCDir {
	for _, e := range d.entries {
		if e.Name() == name {
			result, _ := e.(*AOCDir)
			return result
		}
	}
	if name == d.name {
		return d
	}
	return nil
}

type AOCFile struct {
	name string
	size int
}

func (f *AOCFile) Name() string {
	return f.name
}

func (f *AOCFile) Size() int {
	return f.size
}

func (f *AOCFile) String() string {
	return fmt.Sprintf("- %s (file, size: %d)", f.name, f.size)
}

func NewFile(name string, size int) *AOCFile {
	return &AOCFile{name, size}
}

type Parser struct {
	//processed  []string
	currentDir *AOCDir
}

func (p *Parser) Parse(line string) {
	//p.processed = append(p.processed, line)
	switch {
	case cdCommandPattern.MatchString(line):
		matches := cdCommandPattern.FindStringSubmatch(line)
		nameIndex := cdCommandPattern.SubexpIndex("name")
		name := matches[nameIndex]
		if name == ".." {
			log.Printf("switching to parent dir\n")
			p.currentDir = p.currentDir.parent
			break
		}
		dir := p.currentDir.find(name)
		if dir == nil {
			log.Fatalf("FATAL: Could not find dir %v", name)
		}
		p.currentDir = dir
		log.Printf("switching dir => %v\n", name)

	case lsCommandPattern.MatchString(line):
		log.Printf("Listing directory => %v", p.currentDir.name)

	case dirOutputLinePattern.MatchString(line):
		matches := dirOutputLinePattern.FindStringSubmatch(line)
		nameIndex := dirOutputLinePattern.SubexpIndex("name")
		name := matches[nameIndex]
		dir := NewDir(name, p.currentDir)
		p.currentDir.AddEntry(dir)
		log.Printf("found dir => %v\n", name)

	case fileOutputLinePattern.MatchString(line):
		matches := fileOutputLinePattern.FindStringSubmatch(line)
		nameIndex := fileOutputLinePattern.SubexpIndex("name")
		sizeIndex := fileOutputLinePattern.SubexpIndex("size")

		name := matches[nameIndex]
		size, _ := strconv.Atoi(matches[sizeIndex])
		file := NewFile(name, size)
		p.currentDir.AddEntry(file)

		log.Printf("found file => %v\n", name)
	}
}

func NewParser(root *AOCDir) *Parser {
	return &Parser{
		currentDir: root,
	}
}

func GenerateListing(d *AOCDir, depth int) {
	prefix := strings.Repeat("  ", depth)
	fmt.Printf("%s%s\n", prefix, d.String())
	for _, e := range d.entries {
		if f, ok := e.(*AOCFile); ok {
			fmt.Printf("%s  %s\n", prefix, f.String())
		} else {
			sd, _ := e.(*AOCDir)
			GenerateListing(sd, depth+1)
		}
	}
}
