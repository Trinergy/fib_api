package badgerstore

import (
	"strconv"
)

// These functions are used for basic math operations with strings

func addStr(s string, s2 string) (string, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(n + n2), nil
}

func strIncrement(s string) (string, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}

	n++
	return strconv.Itoa(n), nil
}

func strDecrement(s string) (string, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}

	n--
	return strconv.Itoa(n), nil
}
