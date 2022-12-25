package day15

import (
	"fmt"
	"testing"
)

func TestOneSensor(t *testing.T) {
    //Sensor at x=8, y=7: closest beacon is at x=2, y=10
    beacon := Beacon{2, 10}
    sensor := Sensor{8, 7, beacon}
    dist := getManhattanDistance(8,7,2,10)
    m := make(map[string]bool)
    setCantBeBeacon(sensor, beacon, dist, m)
    fmt.Printf("len %d\n", len(m))

    _ = getNotBeaconCount(10, m)
}
func TestGetNumber(t *testing.T) {
    s := "=-1"

    n := getNumber(s)
    if n != -1 {
        t.Fatal("not -1")
    }
}
func TestGetRange(t *testing.T) {
    a := 2
    b := 0

    res := getRange(a,b)

    fmt.Printf("%v\n", res)
}
