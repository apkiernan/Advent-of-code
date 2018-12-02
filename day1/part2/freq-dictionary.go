package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./frequencies.txt")
	onError(err)

	sum := 0
	dict := make(map[int]bool)
	stringData := string(data)
	frequencies := strings.Split(stringData, "\n")

	repeated := runFrequencies(sum, frequencies, dict)
	fmt.Println(repeated)
}

func runFrequencies(sum int, frequencies []string, dict map[int]bool) int {

	for idx := 0; idx < len(frequencies); idx++ {
		if frequencies[idx] != "" {
			num, err := strconv.Atoi(frequencies[idx])
			onError(err)
			sum = sum + num
			if dict[sum] {
				return sum
			}

			dict[sum] = true
		}
	}
	return runFrequencies(sum, frequencies, dict)
}

func onError(err error) {
	if err != nil {
		panic(err)
	}
}
