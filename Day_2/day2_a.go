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

	//solve for the final position
	horizontal := 0
	depth := 0

	for _, val := range values {
		res := strings.Split(val, " ")
		if res[0] == "forward" {
			if i, err := strconv.Atoi(res[1]); err == nil {
				horizontal += i
			}
		} else if res[0] == "up" {
			if i, err := strconv.Atoi(res[1]); err == nil {
				depth -= i
			}
		} else if res[0] == "down" {
			if i, err := strconv.Atoi(res[1]); err == nil {
				depth += i
			}
		}

	}
	fmt.Println(horizontal * depth)

}
