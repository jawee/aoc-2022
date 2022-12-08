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
            // fmt.Printf("Calculating for %d, %d\n", i, j)
                if isVisibleToEdge(i, j, trees) {
                    // fmt.Printf("Got true for %d, %d\n", i, j)
                    res++
                } else {
                    // fmt.Printf("Got false for %d, %d\n", i, j)
                }
            }
        }
    }
    fmt.Printf("%d\n", res)
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


func isVisibleToEdge(i int ,j int, trees [][]int) bool {
    // fmt.Printf("%d, %d\n", i, j)
    if i == 0 || i == len(trees)-1 {
        // fmt.Printf("isVisibleToEdge on edge %t\n", true)
        return true
    }

    if j == 0 || j == len(trees[i])-1 {
        // fmt.Printf("isVisibleToEdge on edge %t\n", true)
        return true
    }

    // visited[i][j] = true
    val := trees[i][j]
    var top bool
    var right bool
    var down bool
    var left bool
    // fmt.Printf("recursing for %d,%d for val %d\n", i, j, val)
    // fmt.Printf("%t\n", res)
    if i-1 >= 0 {
    // top = isVisibleToEdgeRec(i-1, j, trees, val, visited)
        top = isVisibleRec(i-1,j,-1,0,val,trees,getVisitedArray(trees))
        // fmt.Printf("top %t\n", top)
    }
    if j+1 < len(trees[0]) {
        // right = isVisibleToEdgeRec(i, j+1, trees, val, getVisitedArray(trees))
        right = isVisibleRec(i,j+1,0,1,val,trees,getVisitedArray(trees))
        // fmt.Printf("right %t\n", right)
    }
    if i+1 < len(trees) {
        // down = isVisibleToEdgeRec(i+1, j, trees, val, getVisitedArray(trees))
        down = isVisibleRec(i+1,j,1,0,val,trees,getVisitedArray(trees))
        // fmt.Printf("down %t\n", down)
    }
    if j-1 >= 0 {
        left = isVisibleRec(i,j-1,0,-1,val,trees,getVisitedArray(trees))
        // left = isVisibleToEdgeRec(i, j-1, trees, val, getVisitedArray(trees)) 
        // fmt.Printf("left %t\n", left)
    }
    res := top || right || down || left
    return res
}

func getVisitedArray(trees [][]int) [][]bool {
    visited := make([][]bool, len(trees))
    for i := range trees {
        visited[i] = make([]bool, len(trees[i]))
    }
    return visited
}

func isVisibleRec(i,j,x,y,val int, trees [][]int, visited [][]bool) bool {
    // fmt.Printf("%d,%d %d %d %d\n",i,j,x,y,val)
    if visited[i][j] {
        // fmt.Printf("visited\n")
        return false 
    }

    visited[i][j] = true    

    if trees[i][j] >= val {
        // fmt.Printf("%d >= %d at %d, %d\n", trees[i][j], val, i,j)
        return false
    }

    if i+x >= len(trees) || j+y >= len(trees[0]) || i+x < 0 || j+y < 0 {
        // fmt.Printf("edge\n")
        return true
    }

    // fmt.Printf("recurse\n")
    return isVisibleRec(i+x, j+y, x, y, val, trees, visited)
}

// func isVisibleToEdgeRec(i int, j int, trees [][]int, val int, visited [][]bool) bool {
//     if i >= len(trees) || j >= len(trees[0]) {
//         fmt.Printf("at edge %d %d %t\n", i, j, true)
//         return true
//     }
//     if i < 0 || j < 0 {
//         fmt.Printf("at edge %d %d %t\n", i, j, true)
//         return true
//     }
//     fmt.Printf("recurse %d, %d, %d val: %d\n", i, j, trees[i][j], val)
//     //at edge
//     if visited[i][j] {
//         return false
//     }
//
//     visited[i][j] = true
//
//     //tree is higher
//     if (trees[i][j] >= val) {
//         fmt.Printf("is higher %d %d %t\n", trees[i][j], val,false)
//         return false
//     }
//
//     top := isVisibleToEdgeRec(i-1, j, trees, val, visited)
//     right := isVisibleToEdgeRec(i, j+1, trees, val, visited)
//     down := isVisibleToEdgeRec(i+1, j, trees, val, visited)
//     left := isVisibleToEdgeRec(i, j-1, trees, val, visited) 
//
//
//     // fmt.Printf("%d, %d, %d %t %t %t %t\n", i, j, val, top, right, down, left)
//     return top || right || down || left
// }

