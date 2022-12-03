package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day3/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        content := scanner.Text()
        first := content[0:len(content)/2]
        second := content[len(content)/2:]

        for _, i := range first {
            if strings.ContainsRune(second, i) {
                sum = sum + getPriorityForRune(i)
                break
            }
        }
    }

    fmt.Printf("%d\n", sum)
}

func getPriorityForRune(i rune) int {
    if i <= 'z' && i >= 'a' {
        return int(i)-96
    }
    return int(i)-38
}

