package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

const DNAs = "ACGT"

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	if !strings.ContainsRune(DNAs, rune(nucleotide)) {
		return 0, errors.New("Invalid nucleotide")
	}
	count := 0
	for _, r := range d {
		if strings.ContainsRune(DNAs, r) {
			if byte(r) == nucleotide {
				count++
			}
		} else {
			return 0, errors.New("Invalid DNA")
		}
	}
	return count, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range d {
		if strings.ContainsRune(DNAs, r) {
			h[byte(r)] += 1
		} else {
			return nil, errors.New("Invalid DNA")
		}
	}
	return h, nil
}
