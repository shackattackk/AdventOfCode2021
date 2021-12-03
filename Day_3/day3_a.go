package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//	File Opening
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("File not found")
		return
	}
	defer file.Close()

	//  Read Input

	sc := bufio.NewScanner(file)
	var values []string
	for sc.Scan() {
		values = append(values, sc.Text())
	}

	hashmap := make(map[int]int) // init empty hashmap

	//build hashmap to check for most common bit
	for _, val := range values {
		for index, digit := range val {
			if string(digit) == "0" {
				hashmap[index] -= 1
			} else if string(digit) == "1" {
				hashmap[index] += 1
			}

		}
	}

	var gamma strings.Builder
	var epsilon strings.Builder

	for i := 0; i < len(hashmap); i++ {
		if hashmap[i] > 0 {
			gamma.WriteString("1")
		} else if hashmap[i] < 0 {
			gamma.WriteString("0")
		}
	}
	for i := 0; i < len(hashmap); i++ {
		if hashmap[i] > 0 {
			epsilon.WriteString("0")
		} else if hashmap[i] < 0 {
			epsilon.WriteString("1")
		}
	}
	var gammarate int64 = 0
	var epsilonrate int64 = 0
	if i, err := strconv.ParseInt(gamma.String(), 2, 64); err != nil {
		fmt.Println(err)
	} else {
		gammarate = i
	}
	if i, err := strconv.ParseInt(epsilon.String(), 2, 64); err != nil {
		fmt.Println(err)
	} else {
		epsilonrate = i
	}

	fmt.Println(gammarate * epsilonrate)
}
