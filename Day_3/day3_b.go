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
	var values []string
	for sc.Scan() {
		values = append(values, sc.Text())
	}

	oxygen := filter(values, 0, true)
	scrubber := filter(values, 0, false)
	fmt.Println(bin2Dec(oxygen) * bin2Dec(scrubber))

}

func getCommonBit(tempArr []string, i int) (uint8, uint8) {
	zerocnt, onecnt := 0, 0

	for _, value := range tempArr {
		if value[i] == '0' {
			zerocnt++
		} else {
			onecnt++
		}
	}

	if zerocnt > onecnt {
		return '0', '1'
	} else {
		return '1', '0'
	}
}

func filter(resArr []string, i int, wildcard bool) string {
	if len(resArr) == 1 {
		return resArr[0]
	}
	big, small := getCommonBit(resArr, i)
	cmp := small
	if wildcard {
		cmp = big
	}
	filteredArr := make([]string, 0)
	for _, char := range resArr {
		if char[i] == cmp {
			filteredArr = append(filteredArr, char)
		}
	}
	return filter(filteredArr, i+1, wildcard)
}

func bin2Dec(s string) (res int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(i)

}
