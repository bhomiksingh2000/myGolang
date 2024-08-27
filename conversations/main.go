package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks", input)

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {

	} else {
		fmt.Println("added 1", rating+1)
	}

}
