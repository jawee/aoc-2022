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
    // fmt.Printf("1,1: true Res: %t\n", isVisibleToEdge(1,1,trees))
    // fmt.Printf("1,2: true Res: %t\n", isVisibleToEdge(1,2,trees))
    // fmt.Printf("1,3: false Res: %t\n", isVisibleToEdge(1,3,trees))
    // fmt.Printf("2,1: true Res: %t\n", isVisibleToEdge(2,1,trees))
    // fmt.Printf("2,2: false Res: %t\n", isVisibleToEdge(2,2,trees))
    // fmt.Printf("2,3: true Res: %t\n", isVisibleToEdge(2,3,trees))
    // fmt.Printf("3,1: false Res: %t\n", isVisibleToEdge(3,1,trees))
    // fmt.Printf("3,2: true Res: %t\n", isVisibleToEdge(3,2,trees))
    // fmt.Printf("3,3: false Res: %t\n", isVisibleToEdge(3,3,trees))
}

func getViewScore(i,j int, trees [][]int) int {
    if i == 0 || i == len(trees)-1 {
        return 1
    }

    if j == 0 || j == len(trees[i])-1 {
        // fmt.Printf("isVisibleToEdge on edge %t\n", true)
        return 1
    }

    // visited[i][j] = true
    val := trees[i][j]
    var top int
    var right int
    var down int
    var left int
    // fmt.Printf("recursing for %d,%d for val %d\n", i, j, val)
    // fmt.Printf("%t\n", res)
    if i-1 >= 0 {
        top = getViewScoreRec(i-1,j,-1,0,val,trees)
        // fmt.Printf("top %t\n", top)
    }
    if j+1 < len(trees[0]) {
        right = getViewScoreRec(i,j+1,0,1,val,trees)
        // fmt.Printf("right %t\n", right)
    }
    if i+1 < len(trees) {
        down = getViewScoreRec(i+1,j,1,0,val,trees)
        // fmt.Printf("down %t\n", down)
    }
    if j-1 >= 0 {
        left = getViewScoreRec(i,j-1,0,-1,val,trees)
        // fmt.Printf("left %t\n", left)
    }

    return top * right * down * left
}

func getViewScoreRec(i,j,x,y,val int, trees [][]int) int {
    // fmt.Printf("%d,%d %d %d %d\n",i,j,x,y,val)

    if trees[i][j] >= val {
        // fmt.Printf("%d >= %d at %d, %d\n", trees[i][j], val, i,j)
        return 1
    }

    if i+x >= len(trees) || j+y >= len(trees[0]) || i+x < 0 || j+y < 0 {
        // fmt.Printf("edge\n")
        return 1
    }

    // fmt.Printf("recurse\n")
    return 1 + getViewScoreRec(i+x, j+y, x, y, val, trees)
}
