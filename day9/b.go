package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day9/testinput.txt")
    // file, err := os.Open(pwd + "/day9/testinput2.txt")
    file, err := os.Open(pwd + "/day9/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    visited := map[Coords]bool{}
    visited[Coords{0,0}] = true
    rope := make([]Coords, 0)
    for i := 0; i < 10; i++ {
        rope = append(rope, Coords{0,0})
    }
    for scanner.Scan() {
        inputs := strings.Split(scanner.Text(), " ")
        n, _ := strconv.Atoi(inputs[1])
        visited, rope = moveRope(inputs[0], n, rope, visited)
    }

    res := len(visited) 
    fmt.Printf("%d\n", res)
}

func isCloseEnough(first Coords, second Coords) bool {
    return math.Abs(float64(first.x-second.x)) < 2 && math.Abs(float64(first.y-second.y)) < 2
}

func moveRope(direction string, n int, rope []Coords, visited map[Coords]bool) (map[Coords]bool, []Coords) {
    x, y := getDirection(direction)
    for i := 0; i < n; i++ {
        rope[0].x += x
        rope[0].y += y
        for j := 1; j < len(rope); j++ {
            currRope := &rope[j]

            if isCloseEnough(rope[j], rope[j-1]) {
                break
            }
            
            prevRope := &rope[j-1]

            if prevRope.y > currRope.y {
                currRope.y++
            }

            if prevRope.y < currRope.y {
                currRope.y--
            }

            if prevRope.x > currRope.x {
                currRope.x++
            }

            if prevRope.x < currRope.x {
                currRope.x--
            }
            
            if j == 9 {
                visited[*currRope] = true
            }
        }
    }
    return visited, rope
}
