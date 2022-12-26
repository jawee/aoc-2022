package day15

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func A() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/day15/testinput.txt")
 //    y := 10
	file, err := os.Open(pwd + "/day15/input.txt")
    y := 2000000

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    sensors := make([]Sensor, 0)
    // m := make(map[string]bool)
	for scanner.Scan() {
		l1 := scanner.Text()
        // l1 := "Sensor at x=8, y=7: closest beacon is at x=2, y=10"
        // fmt.Printf("%s\n", l1)
        l := strings.Split(l1, " ")
        sX := getNumber(l[2])
        sY := getNumber(l[3])
        bX := getNumber(l[8])
        bY := getNumber(l[9])

        dist := getManhattanDistance(sX,sY,bX,bY)
        beacon := Beacon{bX, bY}
        sensor := Sensor{sX, sY, beacon, dist}
        sensors = append(sensors, sensor)

        // fmt.Printf("%d %d %d %d %d\n", sX, sY, bX, bY, dist)
	}

    notBeaconCount := getNotBeaconCount(y, sensors)
    fmt.Printf("%d\n", notBeaconCount)
}

func getNotBeaconCount(y int, sensors []Sensor) int {
    max := 0
    min := 9999999999999999
    for _, s := range sensors {
        if s.x-s.dist < min {
            min = s.x-s.dist
        }
        if s.x+s.dist > max {
            max = s.x+s.dist
        }
    }

    count := 0
    for _, x := range getRange(min,max) {
        b := false
        for _, s := range sensors {
            if s.b.x == x && s.b.y == y {
                b = false
                break
            }
            if getManhattanDistance(s.x,s.y,x,y) <= s.dist {
                b = true
                break
            }
        }
        if b {
            // fmt.Printf("#")
            count++
        } 
    }

    return count
}

func getRange(a,b int) []int {
    r := make([]int, 0)

    for a <= b {
       r = append(r,a)
       a++
    }

    for a >= b {
       r = append(r,a)
       a--
    }

    return r
}

func getManhattanDistance(sx,sy,bx,by int) int {
    return int(math.Abs(float64(sx - bx)) + math.Abs(float64(sy - by)))
}

func getNumber(s string) int {
    found := false
    newStr := ""
    for _,c := range strings.Split(s, "") {
        if found && c != "" && c != "," && c != ":" {
            // fmt.Printf("%s\n", c)
            newStr += c
        }
        if c == "=" {
            found = true
        }
    }
    num, _ := strconv.Atoi(newStr)
    return num
}

type Sensor struct {
    x int
    y int
    b Beacon
    dist int
}

type Beacon struct {
    x int
    y int
}

