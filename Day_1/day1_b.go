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

	l := 0
	max_sum, temp := 0, 0
	result := 0

	for index, val := range values {
		temp += val
		if index-l+1 == 3 {
			if temp > max_sum && l != 0 {
				result++
			}
			max_sum = temp
			temp -= values[l]
			l++
		}
	}
	fmt.Println(result)
}
