package main

import (
	"fmt"
	"github.com/rm511130/factoid/Godeps/_workspace/src/github.com/go-martini/martini"
	"strconv"
)

func fact(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	m := martini.Classic()
	m.Get("/**", func(params martini.Params) string {
		s, err := strconv.Atoi(params["_1"])
		fmt.Sprintf("%s %d\n", err, s)
		return "[Factoid] Calculating Factorials: " + params["_1"] + "! = " + fmt.Sprintf("%d\n", fact(uint64(s)))
	})
	m.Run()
}
