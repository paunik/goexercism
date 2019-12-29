package strain

type Ints []int
type Lists [][]int
type Strings []string

type OperationInt func(x int) bool
type OperationString func(x string) bool
type OperationList func(x []int) bool

func (ins Ints) Keep(op OperationInt) (out Ints) {
	for _, i := range ins {
		if op(i) {
			out = append(out, i)
		}
	}
	return
}

func (ins Ints) Discard(op OperationInt) (out Ints) {
	for _, i := range ins {
		if !op(i) {
			out = append(out, i)
		}
	}
	return
}

func (ins Lists) Discard(op OperationList) (out Lists) {
	for _, i := range ins {
		if !op(i) {
			out = append(out, i)
		}
	}
	return
}

func (ins Lists) Keep(op OperationList) (out Lists) {
	for _, i := range ins {
		if op(i) {
			out = append(out, i)
		}
	}
	return
}

func (ins Strings) Keep(op OperationString) (out Strings) {
	for _, i := range ins {
		if op(i) {
			out = append(out, i)
		}
	}
	return
}

func (ins Strings) Discard(op OperationString) (out Strings) {
	for _, i := range ins {
		if !op(i) {
			out = append(out, i)
		}
	}
	return
}
