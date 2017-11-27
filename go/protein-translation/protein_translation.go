// Package protein provides functions for translating RNA sequences to proteins.
package protein

const testVersion = 1

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns the Amino acid of the given codon
func FromCodon(codon string) string {
	return codonMap[codon]
}

// FromRNA returns the protein of the given rna
func FromRNA(rna string) (proteins []string) {
	for i := 0; i < len(rna); i += 3 {
		aminoAcid := FromCodon(rna[i : i+3])
		if aminoAcid == "STOP" {
			return
		}
		proteins = append(proteins, aminoAcid)
	}
	return
}
