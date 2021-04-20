package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line[0]) > 1 && len(line[1]) > 1 {
			fmt.Println("SET", line[1], line[0])
		}
	}
}

func main() {
	scanner()
}
