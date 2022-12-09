package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
    x int;
    y int;
}

func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day9/testinput.txt")
    file, err := os.Open(pwd + "/day9/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    visited := map[Coords]bool{}
    visited[Coords{0,0}] = true
    head := &Coords{0,0}
    tail := &Coords{0,0}
    for scanner.Scan() {
        inputs := strings.Split(scanner.Text(), " ")
        n, _ := strconv.Atoi(inputs[1])
        visited = move(inputs[0], n, head, tail, visited)
    }

    res := len(visited) 
    fmt.Printf("%d\n", res)
}

func move(direction string, n int, head *Coords, tail *Coords, visited map[Coords]bool) map[Coords]bool {
    for i:= 0; i < n; i++ {
        x, y := getDirection(direction)
        head.x = head.x + x
        head.y = head.y + y
        if (tail.x == head.x-x && tail.y == head.y-y) || getDistance(head, tail) <= 1 {
            continue
        }

        if getDistance(head, tail) > 1 {
            tail.x = head.x - x
            tail.y = head.y - y
            visited[Coords{tail.x, tail.y}] = true
        }
    }

    return visited
}
func getDistance(head *Coords, tail *Coords) int {
    res := math.Sqrt(math.Pow(float64(head.x-tail.x), 2) + math.Pow(float64(head.y-tail.y), 2))
    return int(res)
}
func getDirection(direction string) (x int, y int) {
    switch direction {
    case "R":
        return 1,0
    case "U":
        return 0,1
    case "L":
        return -1, 0
    default:
        return 0,-1
    }
}
