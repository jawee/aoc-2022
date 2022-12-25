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
	file, err := os.Open(pwd + "/day15/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    sensors := make([]Sensor, 0)
    m := make(map[string]bool)
	for scanner.Scan() {
		l1 := scanner.Text()
        // l1 := "Sensor at x=8, y=7: closest beacon is at x=2, y=10"
        // fmt.Printf("%s\n", l1)
        l := strings.Split(l1, " ")
        sX := getNumber(l[2])
        sY := getNumber(l[3])
        bX := getNumber(l[8])
        bY := getNumber(l[9])

        beacon := Beacon{bX, bY}
        sensor := Sensor{sX, sY, beacon}
        dist := getManhattanDistance(sX,sY,bX,bY)
        sensors = append(sensors, sensor)
        setCantBeBeacon(sensor, beacon, dist, m)

        // fmt.Printf("%d %d %d %d %d\n", sX, sY, bX, bY, dist)
	}


    notBeaconCount := getNotBeaconCount(10, m)
    fmt.Printf("%d\n", notBeaconCount)
}

func getNotBeaconCount(y int, m map[string]bool) int {
    max := 0
    min := 9999999999999999

    for k := range m {
        // fmt.Printf("%s\n", k)
        p := strings.Split(k, ",")
        x, _ := strconv.Atoi(p[0])
        if x > max {
            max = x
        }
        if x < min {
            min = x
        }
    }

    // fmt.Printf("%v", m)
    // sum := len(sensors)
	fmt.Printf("%d %d\n", min,max)
    count := 0

    for _, x := range getRange(min,max) {
        s := fmt.Sprintf("%d,%d", x, y)
        // fmt.Printf("%s %v\n", s, m[s])
        if m[s] == true {
            count++
        }
    }
    return count
}



func setCantBeBeacon(s Sensor, b Beacon, dist int, m map[string]bool) {
    fmt.Printf("=============setCantBeBeacon=========================\n")
    xR := getRange(s.x-dist, s.x+dist)
    yR := getRange(s.y-dist, s.y+dist)
    fmt.Printf("xR: %d\n", len(xR))
    fmt.Printf("yR: %d\n", len(yR))
    for _, x := range xR  {
        for _,y := range yR {
            if getManhattanDistance(s.x, s.y, x, y) > dist {
                // fmt.Printf("%d,%d is too far away\n", x,y)
                continue
            }
            if b.x == x && b.y == y {
                // fmt.Printf("%d,%d is equal to %d, %d\n", x,y,b.x,b.y)
            } else {
                // fmt.Printf("Marking true.%d,%d\n", x, y)
                // m[fmt.Sprintf("%d,%d", x, y)] = true
            }
        }
    }
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
}

type Beacon struct {
    x int
    y int
}

