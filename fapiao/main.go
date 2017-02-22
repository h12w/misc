package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type solution struct {
	items []float64
	sum   float64
}

type counter struct {
	a    []int
	base int
}

func newCounter(size, base int) counter {
	return counter{
		a:    make([]int, size),
		base: base,
	}
}

func (c *counter) inc() bool {
	for i := range c.a {
		c.a[i]++
		if c.a[i] < c.base {
			break
		}
		c.a[i] = 0
		if i == len(c.a)-1 {
			return true // overflow
		}
	}
	return false
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("fapiao [target] [fapiao1,fapiao2,...]")
		return
	}
	fmt.Println("Press Ctrl+C to exit...\n")
	target, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	fapiaos := strings.Split(os.Args[2], ",")
	all := make([]float64, len(fapiaos))
	for i, fapiao := range fapiaos {
		var err error
		all[i], err = strconv.ParseFloat(fapiao, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	// exclude larger Fapiao
	sort.Sort(sort.Reverse(sort.Float64Slice(all)))

	best := solution{sum: fsum(all), items: all}
	if best.sum < target {
		fmt.Println("Not enough Fapiao!")
	}
	for size := 1; size < len(all); size++ {
		cnt := newCounter(size, len(all))
		for !cnt.inc() {
			if !hasDup(cnt.a) {
				sum := sumByIndex(all, cnt.a)
				if sum > target {
					if sum < best.sum {
						best = solution{sum: sum, items: getByIndex(all, cnt.a)}
						fmt.Println(best)
					}
				}
			}
		}
	}
	fmt.Println(best)
}

func hasDup(a []int) bool {
	contains := make(map[int]bool)
	for _, n := range a {
		if contains[n] {
			return true
		}
		contains[n] = true
	}
	return false
}

func sumByIndex(a []float64, indexes []int) float64 {
	sum := 0.0
	for _, index := range indexes {
		sum += a[index]
	}
	return sum
}

func getByIndex(a []float64, indexes []int) []float64 {
	values := make([]float64, len(indexes))
	for i := range indexes {
		values[i] = a[indexes[i]]
	}
	return values
}

func fsum(a []float64) float64 {
	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum
}
