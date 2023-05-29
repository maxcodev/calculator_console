package models

import (
	"strconv"
)

type Calc struct {
	Operator string
	Val1     string
	Val2     string
}

func (Calc) Operate(operator string, val1 string, val2 string) int {
	value1 := parseable(val1)
	value2 := parseable(val2)
	switch operator {
	case "a":
		return value1 + value2
	case "b":
		return value1 - value2
	case "c":
		return value1 * value2
	case "d":
		return value1 / value2
	default:
		return 0
	}
}

func parseable(val string) int {
	parseado1, _ := strconv.Atoi(val)
	return parseado1
}
