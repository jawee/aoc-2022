package day11

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
    n int;
    op string;
    v string;
    items []int;
    test int;
    t int;
    f int;
    count int;
}

func NewMonkey(n int, op string, v string, items []int, test int, t int, f int) *Monkey {
    return &Monkey{
        n:n,
        op:op,
        v:v,
        items:items,
        test:test,
        t:t,
        f:f,
        count: 0,
    }
}

func (m *Monkey) printMonkey() {
    fmt.Printf("Monkey %d\n ", m.n)
    fmt.Printf("    Starting items:")
    for _, r := range m.items {
        fmt.Printf("%d, ", r)
    }
    fmt.Printf("\n")
    fmt.Printf("    Operation: new = old %s %s\n", m.op, m.v)
    fmt.Printf("    Test: divisible by %d\n", m.test)
    fmt.Printf("        If true: throw to monkey %d\n", m.t)
    fmt.Printf("        If true: throw to monkey %d\n", m.f)
}

func (m *Monkey) printCount() {
    fmt.Printf("Monkey %d inspected items %d times.\n", m.n, m.count)
}

func A() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day11/testinput.txt")
    file, err := os.Open(pwd + "/day11/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
        

    scanner := bufio.NewScanner(file)
    monkeys := getMonkeys(scanner)
    
    for i := 0; i < 20; i++ {
        for _,m := range monkeys {
            // m.printMonkey()
            for _, item := range m.items {
                m.count++
                item = calculateNewValue(m.op, m.v, item)
                item = item / 3
                if item % m.test == 0 {
                    monkeys[m.t].items = append(monkeys[m.t].items, item)
                } else {
                    monkeys[m.f].items = append(monkeys[m.f].items, item)
                }
            }
            m.items = make([]int, 0)
        }
    }

    counts := make([]int, 0)
    for _, m := range monkeys {
        // m.printMonkey()
        // m.printCount()
        counts = append(counts, m.count)
    }

    sort.Ints(counts)
    
    fmt.Printf("%d\n", counts[len(counts)-1]*counts[len(counts)-2])

}

func calculateNewValue(op, val string, item int) int {
    res := 1
    switch op {
    case "*":
        if strings.Trim(val, " ") == "old" {
            res = item * item
        } else {
            v, _ := strconv.Atoi(val)
            res = item * v
        }
    case "+":
        if strings.Trim(val, " ") == "old" {
            res = item + item
        } else {
            v, _ := strconv.Atoi(val)
            res = item + v
        }
    }

    return res
}

func getMonkeys(scanner *bufio.Scanner) []*Monkey {
    monkeys := make([]*Monkey, 0)
    for scanner.Scan() {
        line := scanner.Text()
        //monkey no
        n, _ := strconv.Atoi(string(line[7]))

        //starting items
        scanner.Scan()
        line = scanner.Text()
        line = line[17:]
        itemsStr := strings.Split(line, ",")
        items := make([]int, 0)
        for _, c := range itemsStr {
            v, _ := strconv.Atoi(strings.Trim(c, " "))
            items = append(items, v)
        }

        //Operation
        scanner.Scan()
        line = scanner.Text()
        itemsStr = strings.Split(line, " ")
        operation := itemsStr[len(itemsStr)-2]
        v := itemsStr[len(itemsStr)-1]

        //Test
        scanner.Scan()
        line = scanner.Text()
        itemsStr = strings.Split(line, " ")
        test, _ := strconv.Atoi(itemsStr[len(itemsStr)-1])

        //true dest
        scanner.Scan()
        line = scanner.Text()
        itemsStr = strings.Split(line, " ")
        t, _ := strconv.Atoi(itemsStr[len(itemsStr)-1])
        
        //false dest
        scanner.Scan()
        line = scanner.Text()
        itemsStr = strings.Split(line, " ")
        f, _ := strconv.Atoi(itemsStr[len(itemsStr)-1])

        monkey := NewMonkey(n, operation, v, items, test, t, f)
        monkeys = append(monkeys, monkey)

        scanner.Scan()
    }
    return monkeys
}
