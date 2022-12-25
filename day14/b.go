package day14

import (
	"bufio"
	"fmt"
	"os"
)


func B() {
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

    maxY := getMaxY(objs)
    sum := 0
    for dropSandB(500, objs, maxY) {
       sum++
    }
	fmt.Printf("%d\n", sum)
}

func dropSandB(x int, rl map[string]int, maxY int) bool {
    // fmt.Printf("Dropping sand from x: %d\n", x)
    res := dropSandRecB(x, 0, rl, maxY)
    // fmt.Printf("dropSand:%v\n", rl)
    return res
}

func doesExist(x,y int, rl map[string]int, maxY int) bool {
    if y == maxY + 2 {
        return true
    }
    _, exists := rl[fmt.Sprintf("%d,%d", x, y)]
    return exists
}

func dropSandRecB(x, y int, rl map[string]int, maxY int) bool {
    if y == maxY+2 {
        fmt.Printf("%d, %d\n", maxY, y+2)
        fmt.Printf("Hit floor\n")
        return false
    }

    downExists := doesExist(x, y+1, rl, maxY)
    if !downExists {
        return dropSandRecB(x, y+1, rl, maxY)
    }
    leftExists := doesExist(x-1, y+1, rl, maxY)
    if !leftExists {
        return dropSandRecB(x-1, y+1, rl, maxY)
    }
    rightExists := doesExist(x+1, y+1, rl, maxY)
    if !rightExists {
        return dropSandRecB(x+1, y+1, rl, maxY)
    }

    if doesExist(x, y, rl, maxY) {
        return false
    }
    fmt.Printf("Creating sand at %d,%d\n", x, y)
    rl[fmt.Sprintf("%d,%d", x, y)] = 2
    return true
}
