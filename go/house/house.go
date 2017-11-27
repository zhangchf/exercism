package house

import "bytes"

const testVersion = 1

type Brick struct {
	name, verb string
}

var bricks = []Brick{
	{"the house", "Jack built"},
	{"the malt", "lay in"},
	{"the rat", "ate"},
	{"the cat", "killed"},
	{"the dog", "worried"},
	{"the cow with the crumpled horn", "tossed"},
	{"the maiden all forlorn", "milked"},
	{"the man all tattered and torn", "kissed"},
	{"the priest all shaven and shorn", "married"},
	{"the rooster that crowed in the morn", "woke"},
	{"the farmer sowing his corn", "kept"},
	{"the horse and the hound and the horn", "belonged to"},
}

func Verse(line int) string {
	return "This is " + inner(line-1) + "."
}

func inner(line int) string {
	if line > 0 {
		return bricks[line].name + "\nthat " + bricks[line].verb + " " + inner(line-1)
	} else if line == 0 {
		return bricks[line].name + " that " + bricks[line].verb
	} else {
		return ""
	}
}

func Song() string {
	var result bytes.Buffer
	for i := 0; i < len(bricks); i++ {
		if i > 0 {
			result.WriteString("\n\n")
		}
		result.WriteString(Verse(i + 1))
	}
	return result.String()
}
