package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Thingy struct {
	num int
}

type Thingies []Thingy

func (tt Thingies) Print() {
	for _, t := range tt {
		fmt.Println(t.num)
	}
}

func (tt Thingies) Len() int {
	return len(tt)
}

func (tt Thingies) Less(i, j int) bool {
	return tt[i].num < tt[j].num
}

func (tt Thingies) Swap(i, j int) {
	tt[i].num, tt[j].num = tt[j].num, tt[i].num
}

func main() {
	thingies := Thingies{}

	for _, a := range os.Args {
		if num, err := strconv.Atoi(a); err == nil {
			t := Thingy{num}
			thingies = append(thingies, t)
		}
	}
	sort.Sort(thingies)
	thingies.Print()
}
