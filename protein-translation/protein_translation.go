package protein

import "errors"

var proteins = map[string]string{
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
	"UGC": "Cysteine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
}
var stopCodons = map[string]string{
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("stop Codon")
var ErrInvalidBase = errors.New("invalid base")

// FromCodon translates codon into protein, and returns err if some of stop codons are found or some invalid codon code
func FromCodon(input string) (string, error) {
	_, stop := stopCodons[input]
	if stop {
		return "", ErrStop
	}

	p, ok := proteins[input]

	if ok {
		return p, nil
	}

	return "", ErrInvalidBase
}

// FromRNA retruns slice of Proteins based on RNA input, if stop code is found stops processing and returns protens
// found so far
func FromRNA(input string) ([]string, error) {
	var codons []string

	for i := 0; i < len(input); i += 3 {
		protein, err := FromCodon(string(input[i : i+3]))

		if err == ErrStop {
			return codons, nil
		}

		if err == ErrInvalidBase {
			return codons, ErrInvalidBase
		}

		codons = append(codons, protein)
	}

	return codons, nil
}
