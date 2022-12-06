package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day6/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input := strings.Split(scanner.Text(), "")
        posOfMarker := getPosOfMarker(input, 14)
        fmt.Printf("%d\n", posOfMarker)
    }
}
