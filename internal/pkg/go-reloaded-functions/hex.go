package goreloaded

import (
	"strconv"
)

func Hex(str string) string {
	dec, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(dec))
}
