package plan

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

// AddDone adds a new '*' entry, for something that has been started and completed on the same day.
func (p *Plan) AddDone(s string) {
	p.Done = p.append("* ", p.Done, s)
}

// AddCompleted adds a new '+' entry, for something that has been started earlier and has been completed on the same day.
func (p *Plan) AddCompleted(s string) {
	p.Completed = p.append("+ ", p.Completed, s)
}

// AddCanceled adds a new '-' entry, for something that has been canceled and is will not be completed anymore.
func (p *Plan) AddCanceled(s string) {
	p.Canceled = p.append("- ", p.Canceled, s)
}

// AddNote adds a generic line to the plan.
func (p *Plan) AddNote(s string) {
	p.Notes = p.append("", p.Notes, s)
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
