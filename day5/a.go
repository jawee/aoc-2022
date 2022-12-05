package day5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, string) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day5/testinput.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    // count := 0
    for scanner.Scan() {
        if scanner.Text() == "" {
            break
        }
        line := strings.Split(scanner.Text(), "")
        numberOfStacks := (len(line)+1)/4
        fmt.Printf("%d StackCount: %d\n", len(line), numberOfStacks)
        for i, c := range line {
            if c == "[" || c == "]" || c == " " {

            } else {
                fmt.Printf("%d %d  %s\n",i, i/(numberOfStacks+1), c)
            }
            // itemArr := line[i*numberOfStacks:i*numberOfStacks+4]
            // item := strings.Join(itemArr, "")
            // fmt.Printf("%s\n", item)
        }
        break
    }

    // fmt.Printf("%d\n", count)
}
