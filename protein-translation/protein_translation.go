package protein

import (
	"errors"
	"fmt"
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
	proteinMap := map[string]int{}
	protein := []string{}

	for i := 0; i < len(rna); i += 3 {
		codon := fmt.Sprintf("%v%v%v", string(rna[i]), string(rna[i+1]), string(rna[i+2]))

		if _, err := FromCodon(codon); err == ErrInvalidBase {
			return nil, ErrInvalidBase
		} else if err == ErrStop {
			break
		}

		if _, exists := proteinMap[codonMap[codon]]; !exists {
			proteinMap[codonMap[codon]] = 1
			protein = append(protein, codonMap[codon])
		}
	}
	return protein, nil
}

func FromCodon(codon string) (string, error) {
	if _, exists := codonMap[codon]; !exists {
		return "", ErrInvalidBase
	}

	if codonMap[codon] == "STOP" {
		return "", ErrStop
	}

	return codonMap[codon], nil
}
