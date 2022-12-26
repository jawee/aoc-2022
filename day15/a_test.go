package day15

import (
	"fmt"
	"testing"
)

func TestOneSensor(t *testing.T) {
    //Sensor at x=8, y=7: closest beacon is at x=2, y=10
    beacon := Beacon{2, 10}
    dist := getManhattanDistance(8,7,2,10)
    sensor := Sensor{8, 7, beacon,dist}

    res := getNotBeaconCount(10, []Sensor{sensor})
    if res != 12 {
        t.Fatalf("Expected 12, Got %d\n", res)
    }
}
func TestTwoSensor(t *testing.T) {
    //Sensor at x=8, y=7: closest beacon is at x=2, y=10
    beacon := Beacon{2, 10}
    dist := getManhattanDistance(8,7,2,10)
    sensor := Sensor{8, 7, beacon,dist}

    res := getNotBeaconCount(10, []Sensor{sensor})
    if res != 12 {
        t.Fatalf("Expected 12, Got %d\n", res)
    }
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
