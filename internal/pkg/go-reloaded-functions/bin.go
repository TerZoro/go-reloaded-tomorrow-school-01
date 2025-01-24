package goreloaded

import (
	"strconv"
)

func Bin(str string) string {
	dec, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(dec))
}
