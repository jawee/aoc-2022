package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day8/testinput.txt")
    file, err := os.Open(pwd + "/day8/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    trees := make([][]int, 0)
    lineCount := 0
    for scanner.Scan() {
        trees = append(trees, make([]int, 0))
        inputs := strings.Split(scanner.Text(), "")
        for _, c := range inputs {
            val, _ := strconv.Atoi(c)
            trees[lineCount] = append(trees[lineCount], val)
        }
        lineCount++
    }

    res := 0
    for i := range trees {
        for j := range trees[i] {
            if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[0])-1 {
                res++
            } else {
                if isVisibleToEdge(i, j, trees) {
                    res++
                } else {
                }
            }
        }
    }
    fmt.Printf("%d\n", res)
}


func isVisibleToEdge(i int ,j int, trees [][]int) bool {
    if i == 0 || i == len(trees)-1 {
        return true
    }

    if j == 0 || j == len(trees[i])-1 {
        return true
    }

    val := trees[i][j]
    var top bool
    var right bool
    var down bool
    var left bool
    if i-1 >= 0 {
        top = isVisibleRec(i-1,j,-1,0,val,trees)
    }
    if j+1 < len(trees[0]) {
        right = isVisibleRec(i,j+1,0,1,val,trees)
    }
    if i+1 < len(trees) {
        down = isVisibleRec(i+1,j,1,0,val,trees)
    }
    if j-1 >= 0 {
        left = isVisibleRec(i,j-1,0,-1,val,trees)
    }
    res := top || right || down || left
    return res
}

func isVisibleRec(i,j,x,y,val int, trees [][]int) bool {
    if trees[i][j] >= val {
        return false
    }

    if i+x >= len(trees) || j+y >= len(trees[0]) || i+x < 0 || j+y < 0 {
        return true
    }

    return isVisibleRec(i+x, j+y, x, y, val, trees)
}

