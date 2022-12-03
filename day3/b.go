package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day3/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        first := scanner.Text()
        scanner.Scan()
        second := scanner.Text()
        scanner.Scan()
        third := scanner.Text()

        for _, i := range first {
            if strings.ContainsRune(second, i) && strings.ContainsRune(third, i) {
                sum = sum + getPriorityForRune(i)
                break
            }
        }
    }

    fmt.Printf("%d\n", sum)
}

