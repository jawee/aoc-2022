package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day4/testinput.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        pairs := strings.Split(scanner.Text(), ",")
        pair1 := pairs[0]
        pair2 := pairs[1]

        if contains(pair1, pair2) {
            count++
        }
    }

    fmt.Printf("%d\n", count)
}

func contains(a, b string) bool {
    alow, ahigh := getHighLow(a)
    blow, bhigh := getHighLow(b)
    
    if alow <= blow && ahigh >= bhigh {
        return true
    }
    if blow <= alow && bhigh >= ahigh {
        return true
    }

    return false
}

func getHighLow(a string) (int, int) {
    str := strings.Split(a, "-")
    low, _ := strconv.Atoi(str[0])
    high, _ := strconv.Atoi(str[1])
    return low, high
}
