package model

type BoolFlag int

const (
	BoolFlagUnspecified BoolFlag = iota // 0: 不明
	BoolFlagFalse                       // 1: False
	BoolFlagTrue                        // 2: True
)

func (f BoolFlag) String() string {
	switch f {
	case BoolFlagFalse:
		return "False"
	case BoolFlagTrue:
		return "True"
	default:
		return "unspecified"
	}
}
