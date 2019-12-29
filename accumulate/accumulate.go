package accumulate

type Operation func(string) string

// Accumulate does in place apply of operation
func Accumulate(in []string, op Operation) []string {

	for index := 0; index < len(in); index++ {
		in[index] = op(in[index])
	}
	return in
}

// AccumulateImmutable does returns copy that holds values after applying operation
func AccumulateImmutable(in []string, op Operation) (out []string) {

	for _, i := range in {
		out = append(out, op(i))
	}

	return
}
