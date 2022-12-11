package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day10/testinput.txt")
    file, err := os.Open(pwd + "/day10/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    cycle := 0
    x := 1
    totalStr := 0
    str := 0
    for scanner.Scan() {
        inputs := strings.Split(scanner.Text(), " ")
        if inputs[0] == "noop" {
            cycle, str = runCycle(x, cycle)
            totalStr += str
        }
        if inputs[0] == "addx" {
            v, _ := strconv.Atoi(inputs[1])
            cycle, str = runCycle(x, cycle)
            totalStr += str
            cycle, str = runCycle(x, cycle)
            totalStr += str
            x = x+v
        }
    }

    fmt.Printf("%d\n", totalStr)
}

func runCycle(x, cycle int) (int, int) {
    cycle++
    str := 0
    
    if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
        fmt.Printf("signal str %d\n", x*cycle)
        str = x*cycle
    }

    return cycle, str
}

