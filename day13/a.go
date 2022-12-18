package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day13/testinput.txt")
    file, err := os.Open(pwd + "/day13/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
        

    scanner := bufio.NewScanner(file)
    
    pairs := make([]Pair, 0)
    idx := 1
    for scanner.Scan() {
    // scanner.Scan()
        l1 := scanner.Text()
        if l1 == "" {
            scanner.Scan()
            l1 = scanner.Text()
        }
        p1 := getObj(l1)
        scanner.Scan()
        l2 := scanner.Text()
        p2 := getObj(l2)

        pair := Pair{p1,p2,idx}
        pairs = append(pairs, pair)
        fmt.Printf("Pair %d: %s vs %s\n", idx, l1,l2)
        idx++
    }


    sum := 0
    for _, pair := range pairs {
        res := getResult(pair)
        fmt.Printf("Handling pair %d. Got res %d\n", pair.idx, res)
        sum += res
        // pair.l1
        // pair.l2
    }

    fmt.Printf("%d\n", sum)

}

func getResult(pair Pair) int {
    for i, o1 := range pair.l1 {
        if i == len(pair.l2) {
            return 0
        }
        o2 := pair.l2[i]
        // fmt.Printf("%v < %v\n", o1, o2)
        if o1.t == o2.t && o1.t == 1 {
            //compare ints
            res := getIntResult(o1,o2)
            switch res {
            case True:
                return pair.idx
            case False:
                return 0
            default:
                continue
            }
        } else if o1.t == o2.t && o1.t == 2 {
            // compare lists
            res := getListResult(o1,o2)
            switch res {
            case True:
                return pair.idx
            case False:
                return 0
            default:
                continue
            }
        } else if o1.t == 1 && o2.t == 2 {
            // fmt.Printf("convert o1 to list\n")
            left := Obj{list:[]Obj{o1},t:2}
            res := getListResult(left,o2)
            switch res {
            case True:
                return pair.idx
            case False:
                return 0
            default:
                continue
            }
            //compare left is int, right is list
        } else if o1.t == 2 && o2.t == 1 {
            //compare left is list, right is int
            // fmt.Printf("convert o2 to list\n")
            right := Obj{list:[]Obj{o2},t:2}
            res := getListResult(o1,right)
            switch res {
            case True:
                return pair.idx
            case False:
                return 0
            default:
                continue
            }
        }
    }
    return pair.idx
}

type Result int

const (
    True Result = iota
    False
    Continue
)

func getListResult(a Obj, b Obj) Result {
    // fmt.Printf("Compare lists\n")
    for i, oa := range a.list {
        if i == len(b.list) {
            return False
        }
        ob := b.list[i]
        if oa.t == ob.t && oa.t == 1 {
            res := getIntResult(oa,ob)
            switch res {
            case True:
                return True
            case False:
                return False
            default:
                continue
            }
        } else if oa.t == ob.t && oa.t == 2 {
            // compare lists
            res := getListResult(oa,ob)
            switch res {
            case True:
                return True
            case False:
                return False
            default:
                continue
            }
        } else if oa.t == 1 && ob.t == 2 {
            left := Obj{list:[]Obj{oa},t:2}
            res := getListResult(left,ob)
            switch res {
            case True:
                return True
            case False:
                return False
            default:
                continue
            }
            //compare left is int, right is list
        } else if oa.t == 2 && ob.t == 1 {
            //compare left is list, right is int
            right := Obj{list:[]Obj{ob},t:2}
            res := getListResult(oa,right)
            switch res {
            case True:
                return True
            case False:
                return False
            default:
                continue
            }
        }
    }

    return True
}

func getIntResult(a Obj, b Obj) Result {
    // fmt.Printf("Compare ints\n")
    if a.val < b.val {
        return True
    } else if (a.val > b.val) {
        return False
    }
    return Continue
}

type Queue struct {
    queue []string
}

func (q *Queue) Peek() string {
    val := q.queue[0]
    return val
}

func (q *Queue) Pop() string {
    val := q.queue[0]
    q.queue = q.queue[1:]
    return val
}

func (q *Queue) Size() int {
    return len(q.queue)
}

func getObj(l string) []Obj {
    // strArr := strings.Split(l, "")
    // fmt.Printf("%s\n", l[1:len(l)-1])
    // remove first and last character
    strArr := strings.Split(l[1:len(l)-1], "")

    queue := Queue{strArr}

    list := make([]Obj, 0)

    for queue.Size() > 0 {
        s := queue.Pop()
        // fmt.Printf("%s\n", s)
        if s == "[" {
            obj := getListObj(s,&queue)
            list = append(list, obj)
            continue
        }
        if s == "," {
            // fmt.Printf("comma. continue\n")
            continue
        }
        obj := getNumberObj(s, &queue)
        list = append(list, obj)
    }
    // fmt.Printf("\n")

    return list
}

func getNumberObj(s string, queue *Queue) Obj {
    // fmt.Printf("getNUmerObj %s\n", s)
    numStr := s 
    for queue.Size() > 0 && queue.Peek() != "," && queue.Peek() != "]" {
        s = queue.Pop()
        numStr += s
    }
    // queue.Pop()
    val, err := strconv.Atoi(numStr)
    // fmt.Printf("%s\n", numStr)
    if err != nil {
        fmt.Printf("%s\n", err)
        panic("err")
    }
    return Obj{t:1, val:val}
}

func getListObj(s string, queue *Queue) Obj {
    list := make([]Obj, 0)
    for queue.Peek() != "]" {
        s := queue.Pop()
        if s == "[" {
            obj := getListObj(s, queue)
            list = append(list, obj)
            continue
        }
        if s == "," {
            // fmt.Printf("comma. continue\n")
            continue
        }
        obj := getNumberObj(s, queue)
        list = append(list, obj)
    }
    
    // pop ] off the stack
    queue.Pop()

    
    return Obj{t:2, list:list}
}

type Obj struct {
    t int
    list []Obj
    val int
}

type Pair struct {
    l1 []Obj
    l2 []Obj
    idx int
}
