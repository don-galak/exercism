package strand

import "strings"

func ToRNA(dna string) string {
	dnaToRna := map[string]string{"C": "G", "G": "C", "T": "A", "A": "U"}
	ts := strings.Split(dna, "")
	for i, n := range ts {
		ts[i] = dnaToRna[n]
	}
	return strings.Join(ts, "")
}
