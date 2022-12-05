package day5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func (s *stack) Push(v string) {
//     s.s = append(s.s, v)
// }
//
// func (s *stack) Pop() string {
//     l := len(s.s)
//     if l == 0 {
//         fmt.Printf("Cant pop enough\n")
//         panic("")
//         // return ""
//     }
//     res := s.s[l-1]
//     s.s = s.s[:l-1]
//     return res
// }
func (s *stack) PushMultiple(a []string) {
    s.s = append(s.s,a...)
}
func (s *stack) PopMultiple(n int) []string {
    l := len(s.s)
    res := s.s[l-n:]
    s.s = s.s[:l-n]

    return res
}
func B() {
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
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        n,f,t := getValues(line)
        
        fval := stacks[f-1].PopMultiple(n)
        stacks[t-1].PushMultiple(fval)
        // for i := 0; i < n; i++ {
        //     fval := stacks[f-1].Pop()
        //     if fval != "" {
        //         stacks[t-1].Push(fval)
        //     }
        // }
    }
    // fmt.Printf("%d\n", count)
    for _, c := range stacks {
        fmt.Printf("%s", c.Pop())
    }
    fmt.Printf("\n")
}

// func getValues(s []string) (int, int, int) {
//     n, _ := strconv.Atoi(s[1])
//     f, _ := strconv.Atoi(s[3])
//     t, _ := strconv.Atoi(s[5])
//     return n,f,t
// }
