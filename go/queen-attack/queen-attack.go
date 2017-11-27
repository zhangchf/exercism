package queenattack

import (
	"errors"
)

const testVersion = 2

func CanQueenAttack(white, black string) (bool, error) {
	wc, wr, wok := getPosition(white)
	bc, br, bok := getPosition(black)

	if !wok || !bok || (wc == bc && wr == br) {
		return false, errors.New("Invalid input")
	}

	if wc == bc || wr == br || wc-bc == wr-br || wc-bc == br-wr {
		return true, nil
	}

	return false, nil
}

func getPosition(pos string) (col, row int, ok bool) {
	if len(pos) != 2 {
		return -1, -1, false
	}

	col = int(pos[0] - 'a')
	row = int(pos[1] - '1')

	if col >= 0 && col <= 7 && row >= 0 && row <= 7 {
		return col, row, true
	}

	return -1, -1, false
}
