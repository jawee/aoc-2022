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
	// file, err := os.Open(pwd + "/day14/testinput.txt")
	file, err := os.Open(pwd + "/day14/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    objs := make(map[string]int)
    // rocks := make([]Rock, 0)

	for scanner.Scan() {
		l1 := scanner.Text()
        getRocksForInput(l1, objs)
	}

    // sum := len(objs)
    // for k := range objs {
    //     fmt.Println(k)
    // }
    sum := 0
    for dropSand(500, objs) {
       sum++
    }
	fmt.Printf("%d\n", sum)
}

func dropSand(x int, rl map[string]int) bool {
    // fmt.Printf("Dropping sand from x: %d\n", x)
    res := dropSandRec(x, 0, rl)
    // fmt.Printf("dropSand:%v\n", rl)
    return res
}

func dropSandRec(x, y int, rl map[string]int) bool {
    maxY := getMaxY(rl)

    if maxY < y+1 {
        fmt.Printf("Falling to infinity\n")
        return false
    }

    _, downExists := rl[fmt.Sprintf("%d,%d", x, y+1)]
    if !downExists {
        return dropSandRec(x, y+1, rl)
    }
    _, leftExists := rl[fmt.Sprintf("%d,%d", x-1, y+1)]
    if !leftExists {
        return dropSandRec(x-1, y+1, rl)
    }
    _, rightExists := rl[fmt.Sprintf("%d,%d", x+1, y+1)]
    if !rightExists {
        return dropSandRec(x+1, y+1, rl)
    }

    fmt.Printf("Creating sand at %d,%d\n", x, y)
    rl[fmt.Sprintf("%d,%d", x, y)] = 2
    return true
}

func getMaxY(rl map[string]int) int {
    max := 0

    for k := range rl {
        p := strings.Split(k, ",")
        y, _ := strconv.Atoi(p[1])
        if y > max {
            max = y
        }
    }

    return max
}

func getRocksForInput(l1 string, rl map[string]int) {
    segments := strings.Split(l1, "->")
    fx := -1
    fy := -1
    for _, s := range segments {
        coords := strings.Split(strings.Trim(s, " "), ",")
        x, _ := strconv.Atoi(coords[0])
        y, _ := strconv.Atoi(coords[1])
        // fmt.Printf("Handle %d,%d %d,%d\n",x,y, fx,fy)

        if fx == -1 && fy == -1 {
            fx = x
            fy = y
            // rl[fmt.Sprintf("%d,%d", x, y)] = Obj{1}
            // fmt.Printf("continue %d,%d\n",x,y)
            continue
        }

        if x == fx {
            // fmt.Printf("Handling y-movement\n")
            //handle y-movement
            dir := 0
            if y > fy {
                dir = 1
                for y >= fy {
                    // fmt.Printf("Creating rock at %d,%d\n", x,fy)
                    rl[fmt.Sprintf("%d,%d", x, fy)] = 1
                    fy += dir
                }
            } else {
                dir = -1
                for y <= fy {
                    // fmt.Printf("Creating rock at %d,%d\n", x,fy)
                    rl[fmt.Sprintf("%d,%d", x, fy)] = 1
                    fy += dir
                }
            }

            fx = x
            fy = y
        } else if y == fy {
            // fmt.Printf("Handling x-movement\n")
            //handle x-movement
            dir := 0
            if x > fx {
                dir = 1
                for x >= fx {
                    // fmt.Printf("Creating rock at %d,%d\n", fx,y)
                    rl[fmt.Sprintf("%d,%d", fx, y)] = 1 
                    fx += dir
                }
            } else {
                dir = -1
                for x <= fx {
                    // fmt.Printf("Creating rock at %d,%d\n", fx,y)
                    rl[fmt.Sprintf("%d,%d", fx, y)] = 1
                    fx += dir
                }
            }

            fx = x
            fy = y
        }
    }

    // return rl
}
