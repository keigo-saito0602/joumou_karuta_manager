package model

type BoolFlag int

const (
	BoolFlagUnspecified BoolFlag = iota // 0: 不明
	BoolFlagTrue                        // 1: True
	BoolFlagFalse                       // 2: False
	//                         // 2: True
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
