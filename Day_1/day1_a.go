package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var values []int
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		values = append(values, num)
	}
	result := 0

	for index, val := range values {
		if index != 0 {
			if val > values[index-1] {
				result++
			}

		}
	}
	fmt.Println(result)
}
