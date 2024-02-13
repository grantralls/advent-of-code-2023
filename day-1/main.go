package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day-1/input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Solution2(file))
}
