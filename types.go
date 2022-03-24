package deepcolor

type ResultType int

func (t ResultType) Contains(resultType ResultType) bool {
	return (t & resultType) != 0
}

const (
	ResultTypeText ResultType = 0b1
	ResultTypeHTMl ResultType = 0b10
	ResultTypeJson ResultType = 0b100
)
