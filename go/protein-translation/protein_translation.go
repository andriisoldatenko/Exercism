package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

func convert(in string) (string, error) {
	switch in {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromCodon(codon string) (string, error) {
	return convert(codon)
}

func FromRNA(rna string) ([]string, error) {
	results := []string{}
	for i := 0; i < len(rna); i += 3 {
		codon, err := convert(rna[i : i+3])
		if err != nil {
			if errors.Is(err, ErrStop) {
				return results, nil
			}
			return results, err
		}
		results = append(results, codon)
	}
	return results, nil
}
