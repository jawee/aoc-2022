package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day6/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input := strings.Split(scanner.Text(), "")
        posOfMarker := getPosOfMarker(input)
        fmt.Printf("%d\n", posOfMarker)
    }
}

func getPosOfMarker(input []string) int {
    for i := range input {
        if i < 4 {
            continue
        } else {
            slice := input[i-4:i]
            if sliceIsUnique(slice) {
                return i
            }
        }
    }
     return -1
}

func sliceIsUnique(slice []string) bool {
    check := make(map[string]bool)

    isUnique := true
    for _, val := range slice {
        if _, ok := check[val]; !ok {
            check[val] = true
        } else {
            isUnique = false
        }
    }
    return isUnique
}
