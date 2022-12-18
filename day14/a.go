package day14

import (
	"bufio"
	"fmt"
	"os"
)

func A() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/day14/testinput.txt")
	// file, err := os.Open(pwd + "/day14/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l1 := scanner.Text()
	}

    sum := 0
	fmt.Printf("%d\n", sum)
}

