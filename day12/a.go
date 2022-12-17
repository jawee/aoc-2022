package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
)

type Node struct {
    coords Coords
    value int
    through *Node
}

type Edge struct {
    node *Node
    weight int
}

type Graph struct {
    Nodes []*Node
    Edges map[Coords][]*Edge
    mutex sync.RWMutex
}

func NewGraph() *Graph {
    return &Graph{ Edges: make(map[Coords][]*Edge),}
}

func (g *Graph) GetNode(coords Coords) *Node {
    g.mutex.RLock()
    defer g.mutex.RUnlock()
    for _, n := range g.Nodes {
        if n.coords.y == coords.y && n.coords.x == coords.x {
            return n
        }
    }
    return nil
}

func (g *Graph) AddNode(n *Node) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    g.Nodes = append(g.Nodes, n)
}

func AddNodes(g *Graph, coords []Coords) (nodes map[Coords]*Node) {
    nodes = make(map[Coords]*Node)
    for _, c := range coords {
        n := &Node{c, math.MaxInt, nil}
        g.AddNode(n)
        nodes[c] = n
    }

    return
}

func (g *Graph) AddEdge(n1,n2 *Node, weight int) {
    g.mutex.Lock()
    defer g.mutex.Unlock()
    g.Edges[n1.coords] = append(g.Edges[n1.coords], &Edge{n2, weight})
    // g.Edges[n2.coords] = append(g.Edges[n2.coords], &Edge{n1, weight})
}


func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day12/testinput.txt")
    file, err := os.Open(pwd + "/day12/input.txt")

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

    nodesList := make([]Coords, 0)
    graph := NewGraph()
    //create grid
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            nodesList = append(nodesList, Coords{i,j})
            // AddNodes(graph, Coords{i,j})
            if grid[i][j] == "S" {
                start = Coords{i,j}
                // grid[i][j]="a"
            }
            if grid[i][j] == "E" {
                end = Coords{i,j}
                // grid[i][j]="z"
            }
        }
    }
    fmt.Printf("%dx%d\n", len(grid), len(grid[0]))

    nodes := AddNodes(graph, nodesList)
    //add edges to graph
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            // fmt.Printf("%s", grid[i][j])
            curr := Coords{i,j}
            currH := getHeight(grid[i][j])
            //top
            if i > 0 {
                if getHeight(grid[i-1][j]) - currH <= 1  { // || getHeight(grid[i-1][j]) - currH == 0  {
                    graph.AddEdge(nodes[curr], nodes[Coords{i-1,j}], 1)
                }
            }
            //right
            if j < len(grid[i])-1 {
                if getHeight(grid[i][j+1]) - currH <= 1 { //|| getHeight(grid[i][j+1]) - currH == 0  {
                    graph.AddEdge(nodes[curr], nodes[Coords{i,j+1}], 1)
                }
            }
            //down
            if i < len(grid)-1 {
                if getHeight(grid[i+1][j]) - currH <= 1 { // || getHeight(grid[i+1][j]) - currH == 0  {
                    graph.AddEdge(nodes[curr], nodes[Coords{i+1,j}], 1)
                }
            }
            //left
            if j > 0 {
                if getHeight(grid[i][j-1]) - currH <= 1 { //|| getHeight(grid[i][j-1]) - currH == 0  {
                    graph.AddEdge(nodes[curr], nodes[Coords{i,j-1}], 1)
                }
            }
        }
        // fmt.Printf("\n")
    }
    // for _, e := range graph.Edges[end] {
    //     fmt.Printf("%v\n", &e.node.coords)
    // }
    dijkstra(graph, start, end)
    for _, node := range graph.Nodes {
        if node.coords == end {
            if node.value == math.MaxInt {
                fmt.Printf("is max int\n")
            }
            fmt.Printf("Shortest path from %v to %v is %d\n", start, end, node.value)
        }
    }
    fmt.Printf("%v%v\n", start, end)
}

func dijkstra(graph *Graph, c Coords, end Coords) {
	visited := make(map[Coords]bool)
	heap := &Heap{}

	startNode := graph.GetNode(c)
	startNode.value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
        if current.coords == end {
            break;
        }
		visited[current.coords] = true
		edges := graph.Edges[current.coords]
		for _, edge := range edges {
            // fmt.Printf("Calculating\n")
			if !visited[edge.node.coords] {
				heap.Push(edge.node)
				if current.value+edge.weight < edge.node.value {
					edge.node.value = current.value + edge.weight
					edge.node.through = current
				}
			}
		}
	}
}
 
type Heap struct {
    mutex sync.RWMutex
    elements []*Node
}

func (h *Heap) Size() int {
    h.mutex.RLock()
    defer h.mutex.RUnlock()
    return len(h.elements)
}
// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Node) {
    h.mutex.Lock()
    defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min value
func (h *Heap) Pop() (i *Node) {
    h.mutex.Lock()
    defer h.mutex.Unlock()
    i = h.elements[0]
    h.elements[0] = h.elements[len(h.elements)-1]
    h.elements = h.elements[:len(h.elements)-1]
    h.rearrange(0)
    return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].value < h.elements[smallest].value {
		smallest = left
	}
	if right < size && h.elements[right].value < h.elements[smallest].value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func getHeight(s string) int {
    if s == "S" {
        s = "a"
    }
    if s == "E" {
        s = "z"
    }
    i := rune(s[0])
    res := int(i)-96
    // fmt.Printf("%s = %d\n", s, res)
    return res
}

type Coords struct {
    x int;
    y int
}
