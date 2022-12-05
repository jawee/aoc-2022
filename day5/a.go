package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
    s []string
}

func (s *stack) Push(v string) {
    s.s = append(s.s, v)
}

func (s *stack) Pop() string {
    l := len(s.s)
    if l == 0 {
        fmt.Printf("Cant pop enough\n")
        panic("")
        // return ""
    }
    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res
}
func (s *stack) Print(n int) {
    fmt.Printf("%d: ", n)
    for _, a := range s.s {
        fmt.Printf("%s", a)
    }
    fmt.Printf("\n")
}
func (s *stack) Reverse() {
    for i, j := 0, len(s.s)-1; i < j; i, j = i+1, j-1 {
        s.s[i], s.s[j] = s.s[j], s.s[i]
    }
}
func NewStack() *stack {
    return &stack { make([]string,0)}
}
func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day5/testinput.txt")
    file, err := os.Open(pwd + "/day5/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    initialized := false
    var stacks []*stack

    //create stacks
    for scanner.Scan() {
        if scanner.Text() == "" {
            break
        }
        line := strings.Split(scanner.Text(), "")
        numberOfStacks := (len(line)+1)/4
        if !initialized {
            initialized = true
            stacks = make([]*stack, numberOfStacks)
            for i := 0; i < numberOfStacks; i++ {
                stacks[i] = NewStack()
            }
        }
        for i, c := range line {
            if c == "1" {
                break
            } else if c == "[" || c == "]" || c == " " {

            } else {
                stacks[i/4].Push(c)
            }
        }
    }
    
    
    for i, c := range stacks {
        c.Reverse()
        c.Print(i+1)
    }
    // move 1 from 2 to 1
    count := 0
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        n,f,t := getValues(line)
        
        for i := 0; i < n; i++ {
            fval := stacks[f-1].Pop()
            if fval != "" {
                stacks[t-1].Push(fval)
            }
        }
        count++
    }
    // fmt.Printf("%d\n", count)
    for _, c := range stacks {
        fmt.Printf("%s", c.Pop())
    }
    fmt.Printf("\n")
}

func getValues(s []string) (int, int, int) {
    n, _ := strconv.Atoi(s[1])
    f, _ := strconv.Atoi(s[3])
    t, _ := strconv.Atoi(s[5])
    return n,f,t
}
