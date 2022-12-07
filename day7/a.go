package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
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
        // input := strings.Split(scanner.Text(), "")
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

    limit := 100000
    res := getSumRecurse(root, limit)

    fmt.Printf("%d\n", res)
}

func getSumRecurse(node *DirNode, limit int) int {
    size := 0

    if node.size <= limit {
        size += node.size
    }

    for _, n := range node.children {
        size += getSumRecurse(n, limit)
    }
    return size
}

func calculateSizes(node *DirNode) int {
    if node == nil {
        return 0
    }
    
    size := 0
    for _, c := range node.files {
        size += c.size
    }

    for _, d := range node.children {
        size += calculateSizes(d)
    }

    node.size = size
    return size
}

func findDirNode(slice []*DirNode, name string) int { 
    for i := range slice {
        if slice[i].name == name {
            return i
        }
    }
    return -1
}

func getCommand(str string) string {
    return str[2:4]
}
func isCommand(str string) bool {
    return str[0] == '$'
}

func CreateDirNode(name string, parent *DirNode) *DirNode {
    return &DirNode{
        name: name,
        children: make([]*DirNode, 0),
        files: make([]*File, 0),
        parent: parent,
    }
}

type DirNode struct {
    name string
    children []*DirNode
    files []*File
    parent *DirNode
    size int
}

type File struct {
    name string
    size int
}

