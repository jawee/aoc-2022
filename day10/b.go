package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CrtLine struct {
    pos int;
    crtLine []string;
    cycle int;
}

func (c *CrtLine) ExecuteCommand(input string) {
    inputs := strings.Split(input, " ")
    if inputs[0] == "noop" {
        c.executeCycle()
    }

    if inputs[0] == "addx" {
        v, _ := strconv.Atoi(inputs[1])
        c.executeCycle()
        c.executeCycle()
        c.pos += v
    }
}

func (c *CrtLine) isLit() bool {
    return c.cycle <= c.pos+1 && c.cycle >= c.pos-1 
}

func (c *CrtLine) executeCycle() {
    if c.isLit() {
        c.crtLine[c.cycle] = "#"
    } else {
        c.crtLine[c.cycle] = "."
    }
    
    c.cycle++

    if c.cycle == 40 {
        c.print()
        c.reset()
    }
}
 
func (c *CrtLine) reset() {
    c.crtLine = make([]string, 40)
    c.cycle = 0
}

func (c *CrtLine) print() {
    for _, s := range c.crtLine {
        fmt.Printf("%s", s)
    }
    fmt.Printf("\n")
}

func B() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day10/testinput.txt")
    file, err := os.Open(pwd + "/day10/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    crtLine := CrtLine{pos:1, cycle: 0, crtLine: make([]string, 40)}
    for scanner.Scan() {
        crtLine.ExecuteCommand(scanner.Text())
    }
}
