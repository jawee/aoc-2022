package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day4/input.txt")

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

        if overlaps(pair1, pair2) {
            count++
        }
    }

    fmt.Printf("%d\n", count)
}

func overlaps(a, b string) bool {
    alow, ahigh := getHighLow(a)
    blow, bhigh := getHighLow(b)
    
    if bhigh < alow {
        return false
    }
    
    if blow > ahigh {
        return false
    }

    return true
}
