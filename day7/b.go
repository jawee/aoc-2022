
package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day7/testinput.txt")
    file, err := os.Open(pwd + "/day7/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    root := CreateDirNode("/", nil)
    var curr *DirNode
    for scanner.Scan() {
        line := scanner.Text()
        if isCommand(line) {
            command := getCommand(line)
            if command == "cd" {
                arg := line[5:]
                if arg == "/" {
                    curr = root
                } else if arg == ".." {
                    curr = curr.parent
                } else {
                    idx := findDirNode(curr.children, arg)
                    curr = curr.children[idx]
                }
            } 
        } else {
            //ls output
            if line[:3] == "dir" {
                idx := findDirNode(curr.children, line[4:])
                if idx == -1 {
                    dir := CreateDirNode(line[4:], curr)
                    curr.children = append(curr.children, dir)
                } 
            } else {
                split := strings.Split(line, " ")
                size, _ := strconv.Atoi(split[0])
                file := &File{name: split[1], size: size}
                curr.files = append(curr.files, file)
            }
        }
    }

    _ = calculateSizes(root)

    freeSpace := 70000000-root.size
    limit := 30000000-freeSpace
    res := getSmallestDirectoryLargerThanLimit(root, limit)

    fmt.Printf("%d\n", res)
}

func getSmallestDirectoryLargerThanLimit(node *DirNode, limit int) int {
    size := 0

    if node.size >= limit {
        size = node.size
    }
    
    if node.size < limit {
        return int(^uint(0) >> 1)
    }

    for _, n := range node.children {
        dirSize := getSmallestDirectoryLargerThanLimit(n, limit)
        if dirSize < size {
            size = dirSize
        }
    }
    return size
}

