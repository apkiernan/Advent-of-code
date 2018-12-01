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

	stringData := string(data)
	frequencies := strings.Split(stringData, "\n")

	sum := 0
	for idx := 0; idx < len(frequencies); idx++ {
		if frequencies[idx] != "" {
			num, err := strconv.Atoi(frequencies[idx])
			onError(err)
			sum += num
		}
	}
	fmt.Println(sum)

}

func onError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
