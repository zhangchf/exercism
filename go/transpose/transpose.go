package transpose

import "strings"

func Transpose(in []string) []string {
	out := []string{}

	rows := len(in)
	cols := 0
	for _, s := range in {
		if len(s) > cols {
			cols = len(s)
		}
	}

	for r := 0; r < cols; r++ {
		row := ""
		shouldTrim := false
		for i := 0; i < rows; i++ {
			if r < len(in[i]) {
				row += in[i][r : r+1]
				shouldTrim = false
			} else {
				row += " "
				shouldTrim = true
			}
		}
		if shouldTrim {
			row = strings.TrimRight(row, " ")
		}
		out = append(out, row)
	}
	return out
}
