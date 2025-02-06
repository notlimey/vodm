package arguments

import (
	"fmt"
	"strconv"
	"strings"
)

type Flags struct {
	Concurrent bool
	Output     string
	Limit      int
}

func ParseFlag(arg string, nextArg string, flags *Flags) (bool, error) {
	// Remove the "-" or "--" prefix
	flag := strings.TrimLeft(arg, "-")

	switch flag {
	case "c", "concurrent":
		flags.Concurrent = true
		return false, nil

	case "o", "output":
		if nextArg != "" && !IsFlag(nextArg) {
			flags.Output = nextArg
			return true, nil
		}
		return false, fmt.Errorf("output flag requires a value")

	case "l", "limit":
		if nextArg != "" && !IsFlag(nextArg) {
			limit, err := strconv.Atoi(nextArg)
			if err != nil {
				return false, fmt.Errorf("invalid limit value: %s", nextArg)
			}
			flags.Limit = limit
			return true, nil
		}
		return false, fmt.Errorf("limit flag requires a numeric value")

	default:
		return false, fmt.Errorf("unknown flag: %s", flag)
	}
}

func IsFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}
