package kindergarten

import (
	"errors"
	"strings"
	"sort"
	"regexp"
)

var plantsMap = map[byte]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

type Garden struct {
	plants   []string
	children []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	numOfChild := len(children)
	for i := 0; i < numOfChild; i++ {
		for j := i+1; j < numOfChild; j++ {
			if children[i] == children[j] {
				return nil, errors.New("Invalid children")
			}
		}
	}
	
	if matched, err := regexp.MatchString(`\n[RCGV]+\n[RCGV]+`, diagram); err != nil || !matched {
		return nil, errors.New("Invalid diagram")
	}
	diagram = strings.TrimLeft(diagram,"\n")
	lines := strings.Split(diagram, "\n")
	if numOfChild == 0 || len(lines) != 2 {
		return nil, errors.New("Invalid diagram")
	}
	for _, line := range lines {
		if len(line) != 2*numOfChild {
			return nil, errors.New("Invalid diagram")
		}
	}
	
	sortedChildren := make([]string, numOfChild)
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)
	return &Garden{lines, sortedChildren}, nil
}

func (garden *Garden) Plants(child string) ([]string, bool) {
	index := garden.find(child)
	if index == -1 {
		return nil, false
	}
	plants := garden.plants[0][2*index:2*index+2] + garden.plants[1][2*index:2*index+2]
	
	plantsDetail := make([]string, 4)
	for i, plant := range plants {
		plantsDetail[i] = plantsMap[byte(plant)]
	}
	return plantsDetail, true
}

// find searches for the given child in the garden's children, and return its index,
// if the child is not found, return -1.
func (garden *Garden) find(child string) int {
	for i, c := range garden.children {
		if c == child {
			return i
		}
	}
	return -1
}
