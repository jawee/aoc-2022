package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

    rocks := make([]Rock, 0)
	for scanner.Scan() {
		l1 := scanner.Text()
        rl := getRocksForInput(l1)
        rocks = append(rocks, rl...)
	}

    sum := len(rocks)
	fmt.Printf("%d\n", sum)
}

func getRocksForInput(l1 string) []Rock {
    rl := make([]Rock, 0)
    segments := strings.Split(l1, "->")
    fx := -1
    fy := -1
    for _, s := range segments {
        // fmt.Printf("%s\n", s)
        coords := strings.Split(strings.Trim(s, " "), ",")
        x, _ := strconv.Atoi(coords[0])
        y, _ := strconv.Atoi(coords[1])

        if fx == -1 && fy == -1 {
            fx = x
            fy = y
            continue
        }

        if x == fx {
            //handle y-movement
            dir := 0
            if y > fy {
                dir = 1
            } else {
                dir = -1
            }

            for y != fy {
                rl = append(rl, Rock{x,fy})
                fy += dir
            }
        } else if y == fy {
            //handle x-movement
            dir := 0
            if x > fx {
                dir = 1
            } else {
                dir = -1
            }

            for x != fx {
                rl = append(rl, Rock{fx,y})
                fx += dir
            }
        }
    }

    return rl
}

type Rock struct {
    x int
    y int
}
