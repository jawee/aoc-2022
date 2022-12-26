package day15

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func B() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day15/testinput.txt")
    // maxX := 20
    // maxY := 20
    file, err := os.Open(pwd + "/day15/input.txt")
    maxX := 4000000
    maxY := 4000000

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

    }
    freq := getFreq(maxX, maxY, sensors)

    fmt.Printf("%d\n", freq)
}

func getFreq(maxX, maxY int, sensors []Sensor) int {
    for _, s := range sensors {
        miX := s.x-s.dist-1
        maX := s.x+s.dist+1
        y := s.y

        for i := 0; i <= maX; i++ {
            x := miX + i
            if x < 0 {
                continue
            }
            if x > maxX {
                break
            }

            yp := y + i
            ym := y - i

            found := true

            if yp >= 0 && yp <= maxY {
                for _, s1 := range sensors {
                    if getManhattanDistance(s1.x, s1.y, x, yp) <= s1.dist {
                        found = false
                        break
                    }
                }

                if found {
                    fmt.Printf("%d,%d found\n", x, yp)
                    return (x*4000000)+yp
                }
            }

            if ym >= 0 && ym <= maxY {
                for _, s1 := range sensors {
                    if getManhattanDistance(s1.x, s1.y, x, ym) <= s1.dist {
                        found = false
                        break;
                    }
                }
                if found {
                    fmt.Printf("%d,%d found\n", x, ym)
                    return (x*4000000)+ym
                }
            }
        }
    }

    return -1
}

