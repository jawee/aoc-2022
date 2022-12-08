package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func B() {
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

    max := 0
    for i := range trees {
        for j := range trees[i] {
            max = int(math.Max(float64(max), float64(getViewScore(i, j, trees))))
        }
    }
    fmt.Printf("%d\n", max)
}

func getViewScore(i,j int, trees [][]int) int {
    if i == 0 || i == len(trees)-1 {
        return 1
    }

    if j == 0 || j == len(trees[i])-1 {
        return 1
    }

    val := trees[i][j]
    var top int
    var right int
    var down int
    var left int

    if i-1 >= 0 {
        top = getViewScoreRec(i-1,j,-1,0,val,trees)
    }
    if j+1 < len(trees[0]) {
        right = getViewScoreRec(i,j+1,0,1,val,trees)
    }
    if i+1 < len(trees) {
        down = getViewScoreRec(i+1,j,1,0,val,trees)
    }
    if j-1 >= 0 {
        left = getViewScoreRec(i,j-1,0,-1,val,trees)
    }

    return top * right * down * left
}

func getViewScoreRec(i,j,x,y,val int, trees [][]int) int {

    if trees[i][j] >= val {
        return 1
    }

    if i+x >= len(trees) || j+y >= len(trees[0]) || i+x < 0 || j+y < 0 {
        return 1
    }

    return 1 + getViewScoreRec(i+x, j+y, x, y, val, trees)
}
