package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day12/testinput.txt")
    // file, err := os.Open(pwd + "/day12/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
        

    scanner := bufio.NewScanner(file)
    
    grid := make([][]string, 0)

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), "")
        grid = append(grid, line)
    }

    var start Coords
    var end Coords

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == "S" {
                start = Coords{i,j}
                grid[i][j]="a"
            }
            if grid[i][j] == "E" {
                end = Coords{i,j}
                grid[i][j]="z"
            }
        }
    }
    fmt.Printf("%d\n", findShortestPath(start, end, grid))
}
 
func findShortestPath(start, end Coords, grid [][]string) int {
    // fmt.Printf("%v %v\n", start, end)

    visited := map[Coords]bool{}
    visited[start] = true

    fmt.Printf("%d\n", len(visited))
    fmt.Printf("%v\n", visited[end])

    h := getHeight("a")

    top, topV := findShortestPathRec(Coords{start.x, start.y-1}, end, h, grid, visited)
    right, rightV := findShortestPathRec(Coords{start.x+1, start.y}, end, h, grid, visited)
    down, downV := findShortestPathRec(Coords{start.x, start.y+1}, end, h, grid, visited)
    left, leftV := findShortestPathRec(Coords{start.x-1, start.y}, end, h, grid, visited)
    min := int(math.Min(math.Min(float64(top), float64(right)), math.Min(float64(down), float64(left))))

    if min == top {
        printPath(topV, grid)
    } else if min == right {
        printPath(rightV, grid)
    } else if min == down {
        printPath(downV, grid)
    } else if min == left {
        printPath(leftV, grid)
    }
    return min
}

func printPath(m map[Coords]bool, grid [][]string) {
    fmt.Println("===== Path =====")
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            // fmt.Printf("%d,%d\n", i, j)
            // fmt.Printf("%v\n", Coords{i,j})
            // c := Coords{i,j}
            if m[Coords{x:i,y:j}] {
                fmt.Printf("#")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Printf("\n")
    }
    // for key, c := range m {
    //     if c {
    //         fmt.Printf("%v\n",key)
    //     }
    // }
}

func findShortestPathRec(a, end Coords, prevHeight int, grid [][]string, visited map[Coords]bool) (int, map[Coords]bool) {
    if a == end {
        // fmt.Printf("Found end \n")
        visited[a] = true
        return 1, visited
    }

    if a.x >= len(grid) || a.x < 0 {
        return math.MaxInt32, visited
    }
    
    if a.y >= len(grid[0]) || a.y < 0 {
        return math.MaxInt32, visited
    }

    if visited[a] {
        return math.MaxInt32, visited
    }

    h := getHeight(grid[a.x][a.y])
    if h > prevHeight+1 {
        // fmt.Printf("%d > %d\n", h, prevHeight+1)
        return math.MaxInt32, visited
    }

    if h < prevHeight {
        // fmt.Printf("%d < %d\n", h, prevHeight)
        return math.MaxInt32, visited
    }

    visited[a] = true

    top, topV := findShortestPathRec(Coords{a.x, a.y-1}, end, h, grid, visited)
    right, rightV := findShortestPathRec(Coords{a.x+1, a.y}, end, h, grid, visited)
    down, downV := findShortestPathRec(Coords{a.x, a.y+1}, end, h, grid, visited)
    left, leftV := findShortestPathRec(Coords{a.x-1, a.y}, end, h, grid, visited)


    min := int(math.Min(math.Min(float64(top), float64(right)), math.Min(float64(down), float64(left))))
    if min == math.MaxInt32 {
        visited[a] = false
        return min, visited
    }

    if min == top {
        return 1 + top, topV
    }

    if min == right {
        return 1 + right, rightV
    }


    if min == down {
        return 1 + down, downV
    }

    return 1+left, leftV
}

func getHeight(s string) int {
    i := rune(s[0])
    res := int(i)-96
    // fmt.Printf("%s = %d\n", s, res)
    return res
}

type Coords struct {
    x int;
    y int
}
