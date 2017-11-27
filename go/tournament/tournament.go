package tournament

import (
	"io"
	"bufio"
	"strings"
	"errors"
	"fmt"
	"sort"
)

var ErrInvalidInput = errors.New("Invalid Input")

const (
	TXT_WIN = "win"
	TXT_LOSS = "loss"
	TXT_DRAW = "draw"
)

type TeamInfo struct {
	win, loss, draw int
	matches, points int
}

type SortMap struct {
	data map[string]*TeamInfo
	keys []string
}

func Tally(reader io.Reader, buffer io.Writer) error {
	var bufScanner = bufio.NewScanner(reader)

	var teams = make(map[string]*TeamInfo)
	for {
		if !bufScanner.Scan() {
			err := bufScanner.Err()
			if err == nil || strings.Compare(err.Error(), "EOF") == 0 {
				break
			}
			return bufScanner.Err()
		}
		text := bufScanner.Text()
		if len(text) == 0 || strings.Compare(text[:1], "#") == 0{
			continue
		}
		parts := strings.Split(text, ";")
		if len(parts) != 3 {
			return ErrInvalidInput
		}
		matchTeam1 := parts[0]
		matchTeam2 := parts[1]
		matchResult := parts[2]
		team1, ok := teams[matchTeam1]
		if !ok {
			team1 = &TeamInfo{}
		}
		team2, ok := teams[matchTeam2]
		if !ok {
			team2 = &TeamInfo{}
		}
		switch matchResult {
		case TXT_WIN:
			team1.win++
			team2.loss++
		case TXT_DRAW:
			team1.draw++
			team2.draw++
		case TXT_LOSS:
			team1.loss++
			team2.win++
		default:
			return ErrInvalidInput
		}
		teams[matchTeam1] = team1
		teams[matchTeam2] = team2
	}

	keys := make([]string, len(teams))
	i := 0
	for team, info := range teams {
		keys[i] = team
		info.matches = info.win + info.draw + info.loss
		info.points = info.win*3 + info.draw*1
		i++
	}
	sortMap := SortMap{data: teams, keys: keys}
	sort.Sort(sortMap)

	teamStr := fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	buffer.Write([]byte(teamStr))
	for _, team := range sortMap.keys {
		info := sortMap.data[team]
		teamStr := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", team, info.matches, info.win, info.draw, info.loss, info.points)
		buffer.Write([]byte(teamStr))
	}

	fmt.Println(buffer)
	return nil
}

func (sm SortMap) Len() int {
	return len(sm.keys)
}

func (sm SortMap) Less(i, j int) bool {
	key1, key2 := sm.keys[i], sm.keys[j]
	info1, info2 := sm.data[key1], sm.data[key2]
	if info1.points != info2.points {
		return sm.data[sm.keys[i]].points > sm.data[sm.keys[j]].points
	} else {
		return strings.Compare(key1, key2) < 0
	}
}

func (sm SortMap) Swap(i, j int) {
	sm.keys[i], sm.keys[j] = sm.keys[j], sm.keys[i]
}
