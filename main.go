/*
 * Created on Fri May 16 2025
 *
 * Copyright (c) Rizki Hadiaturrasyid
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	baseUrl := "https://somewebsite/dir/%v.something"
	from := 0        // start
	to := 4          // end
	isPadded := true // use padding for example 001, 002
	padDigit := 4    // padding digit for example 4 means 0001, 0021

	inputs := []string{}
	for i := from; i <= to; i++ {
		if isPadded {
			padded := pad(i, padDigit)
			input := fmt.Sprintf(baseUrl, padded)
			inputs = append(inputs, input)
		} else {
			input := fmt.Sprintf(baseUrl, i)
			inputs = append(inputs, input)
		}
	}

	printToFile(inputs)

}

func pad(input, digit int) string {
	res := strconv.Itoa(input)
	if len(res) < digit {
		required := digit - len(res)
		for i := 0; i < required; i++ {
			res = "0" + res
		}
	}
	return res
}

func printToFile(inputs []string) {
	now := time.Now().Format("20060102150405")
	file, err := os.Create(now + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, input := range inputs {
		_, err = fmt.Fprintln(file, input)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
