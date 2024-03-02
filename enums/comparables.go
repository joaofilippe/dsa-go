package enums

type Comparable int64

const (
	LESS_THAN Comparable = iota
	EQUAL
	GREATER_THAN
)

// NewComparable returns a Comparable based on the given int64
func NewComparable(a int64) Comparable {
	switch a {
	case 0:
		return LESS_THAN
	case 1:
		return EQUAL
	case 2:
		return GREATER_THAN
	default:
		return -1
	}
}

func (c Comparable) String() string {
	switch c {
	case LESS_THAN:
		return "LESS_THAN"
	case EQUAL:
		return "EQUAL"
	case GREATER_THAN:
		return "GREATER_THAN"
	}
	return "UNKNOWN"
}

func CompareInt64(a, b int64) Comparable {
	if a < b {
		return LESS_THAN
	} else if a == b {
		return EQUAL
	} else {
		return GREATER_THAN
	}
}

func Less(a, b int64) bool {
	return a < b
}

func Greater(a, b int64) bool {
	return a > b
}

func (c Comparable) Compare(a, b int64) bool {
	switch c {
	case LESS_THAN:
		return a<b
	case EQUAL:
		return a == b
	case GREATER_THAN:
		return a > b
	}
	return false
}
