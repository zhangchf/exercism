package ocr

import "strings"

var recognizeDigit = map[string]string{
	`
 _ 
| |
|_|
   `: "0",
	`
   
  |
  |
   `: "1",
	`
 _ 
 _|
|_ 
   `: "2",
	`
 _ 
 _|
 _|
   `: "3",
	`
   
|_|
  |
   `: "4",
	`
 _ 
|_ 
 _|
   `: "5",
	`
 _ 
|_ 
|_|
   `: "6",
	`
 _ 
  |
  |
   `: "7",
	`
 _ 
|_|
|_|
   `: "8",
	`
 _ 
|_|
 _|
   `: "9"}

func Recognize(in string) (out []string) {
	in = strings.Trim(in, "\n")
	lines := strings.Split(in, "\n")
	if len(lines)%4 != 0 {
		return
	}
	for _, line := range lines {
		if len(line)%3 != 0 {
			return
		}
	}

	for row := 0; row < len(lines); row += 4 {
		var rowStr string
		for col := 0; col < len(lines[row]); col += 3 {
			var chStr string
			for cr := 0; cr < 4; cr++ {
				chStr += "\n" + lines[row+cr][col:col+3]
			}
			ch, ok := recognizeDigit[chStr]
			if !ok {
				ch = "?"
			}
			rowStr += ch
		}
		out = append(out, rowStr)
	}
	return out
}
