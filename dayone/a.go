package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/dayone/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    max := 0
    currSum := 0
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        if scanner.Text() == "" {
            max = int(math.Max((float64(max)), float64(currSum)))
            currSum = 0
        }
        v, _ := strconv.Atoi(scanner.Text())
        currSum = currSum + v
    }
    max = int(math.Max((float64(max)), float64(currSum)))
    fmt.Println(max)
}

