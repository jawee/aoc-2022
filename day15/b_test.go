package day15

import "testing"

func TestGetFreq(t *testing.T) {
    shape := Shape{6,2,5,1}
    beacon := Beacon{4,1}
    dist := getManhattanDistance(4,3,4,1)
    sensor := Sensor{4,3,beacon,dist}

    shapes := []Shape{shape}
    sensors := []Sensor{sensor}

    _ = getFreq(7,7,shapes,sensors);
}

