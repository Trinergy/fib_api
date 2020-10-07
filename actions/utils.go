package actions

import (
	"log"
	"strconv"
)

// These functions are used for basic math operations with strings

func addStr(s string, s2 string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		log.Panic(err)
	}

	return strconv.Itoa(n + n2)
}

func strIncrement(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}

	n++
	return strconv.Itoa(n)
}

// decrement can be smart and never decrement past 0 to handle base cases
func strDecrement(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	n--
	return strconv.Itoa(n)
}
