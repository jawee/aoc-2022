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

		pair := Pair{p1, p2, idx}
		pairs = append(pairs, pair)
		fmt.Printf("Pair %d: %s vs %s\n", idx, l1, l2)
		idx++
	}

	sum := 0
	for _, pair := range pairs {
		res := getResult(pair.o1, pair.o2)
		fmt.Printf("Handling pair %d. Got res %d\n", pair.idx, res)
		if res == True || res == Continue {
			sum += pair.idx
		}
	}

	fmt.Printf("%d\n", sum)
    fmt.Printf("6644 is too high\n")

}

func getResult(l1, l2 Obj) Result {
	fmt.Printf("Compare %s - %s\n", l1.ToString(), l2.ToString())

	if l1.t != l2.t {
		if l1.t == 1 {
			l1 = Obj{list: []Obj{l1}, t: 2}
		} else if l2.t == 1 {
			l2 = Obj{list: []Obj{l2}, t: 2}
		}
	}

	if l1.t == 2 && len(l1.list) == 0 {
        if len(l1.list) == len(l2.list) {
            return Continue
        }
		return True
	}

    idx := 0
	for i, o1 := range l1.list {
        idx = i
		if i == len(l2.list) {
			fmt.Printf("Left %s. Right ran out of items\n", o1.ToString())
			return False
		}
		o2 := l2.list[i]
		fmt.Printf("Compare %s - %s\n", o1.ToString(), o2.ToString())
		if o1.t == o2.t && o1.t == 1 {
			//compare ints
			res := getIntResult(o1, o2)
			switch res {
			case True:
				return True
			case False:
				return False
			default:
				continue
			}
		} else if o1.t == o2.t && o1.t == 2 {
			// compare lists
			fmt.Printf("Compare lists %s - %s\n", o1.ToString(), o2.ToString())
			res := getResult(o1, o2)
			switch res {
			case True:
				return True
			case False:
				return False
			default:
				continue
			}
		} else {
			res := getResult(o1, o2)
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

    if len(l2.list)-1 > idx {
        fmt.Printf("%s vs %s. %d > %d. Return true\n", l1.ToString(), l2.ToString(), len(l2.list)-1, idx)
        return True
    }
	return Continue
}

type Result int

const (
	True Result = iota
	False
	Continue
)

func getIntResult(a Obj, b Obj) Result {
	if a.val < b.val {
		fmt.Printf("Compare ints %d vs %d. Res %s\n", a.val, b.val, "true")
		return True
	} else if a.val > b.val {
		fmt.Printf("Compare ints %d vs %d. Res %s\n", a.val, b.val, "false")
		return False
	}
	fmt.Printf("Compare ints %d vs %d. Res %s\n", a.val, b.val, "continue")
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

func getObj(l string) Obj {
	strArr := strings.Split(l, "")
	// fmt.Printf("%s\n", l[1:len(l)-1])
	// remove first and last character
	// strArr := strings.Split(l[1:len(l)-1], "")

	queue := Queue{strArr}

	var result Obj

	for queue.Size() > 0 {
		s := queue.Pop()
		result = getListObj(s, &queue)
		// // fmt.Printf("%s\n", s)
		// if s == "[" {
		//     obj := getListObj(s,&queue)
		//     list = append(list, obj)
		//     continue
		// }
		// if s == "," {
		//     // fmt.Printf("comma. continue\n")
		//     continue
		// }
		// obj := getNumberObj(s, &queue)
		// list = append(list, obj)
	}
	// fmt.Printf("\n")

	return result
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
	return Obj{t: 1, val: val}
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

	return Obj{t: 2, list: list}
}

type Obj struct {
	t    int
	list []Obj
	val  int
}

func (o *Obj) ToString() string {
	s := ""
	if o.t == 1 {
		return fmt.Sprintf("%d", o.val)
	}
	if o.t == 2 {
		s += "["
		for i, obj := range o.list {
			s += obj.ToString()
			if i != len(o.list)-1 {
				s += ","
			}
		}
		s += "]"
	}
	return s
}

type Pair struct {
	o1  Obj
	o2  Obj
	idx int
}
