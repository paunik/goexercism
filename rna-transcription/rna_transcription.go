package strand

var dnaCompliments = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

/*
 * ToRNA replaces compliment nucleotide
 */
func ToRNA(dna string) string {
	ret := ""

	for _, nucleotide := range dna {
		compliment := dnaCompliments[string(nucleotide)]
		ret += compliment
	}

	return ret
}

// func ToRNA(input string) string {
// 	return strings.NewReplacer(
// 		"G", "C",
// 		"C", "G",
// 		"T", "A",
// 		"A", "U",
// 	).Replace(input)
// }
