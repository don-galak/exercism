package protein

import (
	"errors"
)

var ErrStop = errors.New("error stop")
var ErrInvalidBase = errors.New("error invalid base")

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return nil, ErrInvalidBase
		} else if err == ErrStop {
			break
		}
		result = append(result, protein)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonMap[codon]
	if !exists {
		return "", ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
