package plan

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// EntryType represents a plan file's entry type
type EntryType int8

const (
	// Note is an entry with no prefix
	Note EntryType = iota
	// Done is something completed on the same day.
	Done
	// Completed is something completed on a later day.
	Completed
	// Canceled is something decided against on a later day.
	Canceled
)

// Plan contains all the entries, divided in Done, Completed, Canceled, and Notes.
type Plan struct {
	Done      []string
	Completed []string
	Canceled  []string
	Notes     []string
}

func (p Plan) String() string {
	var b strings.Builder

	writeBlock := func(block []string) {
		if len(block) == 0 {
			return
		}
		for _, line := range block {
			b.WriteString(line)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}
	writeBlock(p.Done)
	writeBlock(p.Completed)
	writeBlock(p.Canceled)
	b.WriteString(strings.TrimSpace(strings.Join(p.Notes, "\n")))
	b.WriteString("\n")

	return b.String()
}

// Add adds a new entry to the plan.
func (p *Plan) Add(entry string, typ EntryType) {
	switch typ {
	case Done:
		p.append("* ", p.Done, entry)
	case Completed:
		p.append("+ ", p.Completed, entry)
	case Canceled:
		p.append("- ", p.Canceled, entry)
	default:
		p.append("", p.Notes, entry)
	}
}

func (p Plan) append(prefix string, b []string, s string) []string {
	return append(b, fmt.Sprintf("%s%s", prefix, s))
}

// Load loads a plan file.
func Load(path string) (Plan, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Plan{}, nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Plan{}, err
	}
	return parse(string(content)), nil
}

func parse(in string) Plan {
	p := Plan{}
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "*") {
			p.Done = append(p.Done, line)
		} else if strings.HasPrefix(line, "+") {
			p.Completed = append(p.Completed, line)
		} else if strings.HasPrefix(line, "-") {
			p.Canceled = append(p.Canceled, line)
		} else {
			p.Notes = append(p.Notes, line)
		}
	}
	return p
}
