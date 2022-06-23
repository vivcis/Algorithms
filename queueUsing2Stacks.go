package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Queue using 2 stacks on hackerrank
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	result := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	inputReader, _, err := reader.ReadLine()
	if err != nil {

		log.Fatal(err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader), "\r\n")))
	if err != nil {
		log.Printf("An error occurred: %v", err.Error())
		return
	}

	for i := 1; i <= n; i++ {
		inputReader, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		if inputReader[0] == '1' {
			val, _ := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader[1:]), "\r\n")))
			result = append(result, val)
		} else if inputReader[0] == '2' {
			result = result[1:]
		} else if inputReader[0] == '3' {
			fmt.Printf("%d\n", result[0])
		}
	}
}
