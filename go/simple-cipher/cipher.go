package cipher

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct{}

type Shift struct {
	displacement int
}

type Vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return new(Caesar)
}

func NewShift(displacement int) Cipher {
	if (displacement >= 1 && displacement <= 25) || (displacement >= -25 && displacement <= -1) {
		return &Shift{displacement}
	}
	return nil
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	isAlla := true
	for _, r := range key {
		if !unicode.IsLower(r) {
			return nil
		}
		if r != 'a' {
			isAlla = false
		}
	}
	if isAlla {
		return nil
	}
	return &Vigenere{key}
}

func (caesar *Caesar) Encode(input string) string {
	input = normalize(input)
	return strings.Map(func(r rune) rune {
		return 'a' + (r-'a'+3)%26
	}, input)
}

func (caesar *Caesar) Decode(input string) string {
	input = normalize(input)
	return strings.Map(func(r rune) rune {
		return 'a' + (r-'a'-3+26)%26
	}, input)
}

func (shift *Shift) Encode(input string) string {
	input = normalize(input)
	return strings.Map(func(r rune) rune {
		return 'a' + (r-'a'+int32(shift.displacement)+26)%26
	}, input)
}

func (shift *Shift) Decode(input string) string {
	input = normalize(input)
	return strings.Map(func(r rune) rune {
		return 'a' + (r-'a'-int32(shift.displacement)+26)%26
	}, input)
}

func (vigenere *Vigenere) Encode(input string) string {
	input = normalize(input)
	key := vigenere.key
	for len(key) < len(input) {
		key += key
	}
	result := []rune{}
	for i, r := range input {
		result = append(result, (r-'a'+rune(key[i]-'a')+26)%26+'a')
	}
	return string(result)
}

func (vigenere *Vigenere) Decode(input string) string {
	input = normalize(input)
	key := vigenere.key
	for len(key) < len(input) {
		key += key
	}
	result := []rune{}
	for i, r := range input {
		result = append(result, (r-'a'-rune(key[i]-'a')+26)%26+'a')
	}
	return string(result)
}

func normalize(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, input)
}
