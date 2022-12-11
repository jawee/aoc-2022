package day11

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BMonkey struct {
    n int;
    op string;
    v string;
    items []*big.Int;
    test int;
    t int;
    f int;
    count int;
}

func NewBMonkey(n int, op string, v string, items []*big.Int, test int, t int, f int) *BMonkey {
    return &BMonkey{
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

func (m *BMonkey) printMonkey() {
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

func (m *BMonkey) printCount() {
    fmt.Printf("Monkey %d inspected items %d times.\n", m.n, m.count)
}

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day11/testinput.txt")
    // file, err := os.Open(pwd + "/day11/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
        

    scanner := bufio.NewScanner(file)
    monkeys := getBMonkeys(scanner)
    
    for i := 0; i < 10000; i++ {
        if i == 1 || i == 20 || i % 1000 == 0 {
            fmt.Printf("== After round %d ==\n", i)
            for _, m := range monkeys {
                m.printCount()
            }
        }
        for _,m := range monkeys {
            // m.printMonkey()
            for _, item := range m.items {
                m.count++
                item = calculateNewBValue(m.op, m.v, item)
                res := new(big.Int)
                res.Mod(item, big.NewInt(int64(m.test)))
                if res.Cmp(big.NewInt(0)) == 0 {
                    monkeys[m.t].items = append(monkeys[m.t].items, item)
                } else {
                    monkeys[m.f].items = append(monkeys[m.f].items, item)
                }
            }
            m.items = make([]*big.Int, 0)
            // fmt.Printf("calculating %d\n", i)
        }
    }

    counts := make([]int, 0)
    fmt.Printf("== After round 10000 ==\n")
    for _, m := range monkeys {
        // m.printMonkey()
        m.printCount()
        counts = append(counts, m.count)
    }

    sort.Ints(counts)
    
    res := counts[len(counts)-1]*counts[len(counts)-2]

    fmt.Printf("%d\n", res)

}

func calculateNewBValue(op, val string, item *big.Int) *big.Int {
    b := new(big.Int)
    switch op {
    case "*":
        if strings.Trim(val, " ") == "old" {
            // fmt.Printf("multiply\n")
            fmt.Printf("item is now %d long\n", len(item.String()))
            b.Exp(item, big.NewInt(2), nil)
            // fmt.Printf("done multiply\n")
        } else {
            // fmt.Printf("strconv\n")
            v, _ := strconv.Atoi(val)
            // fmt.Printf("multiply\n")
            b.Mul(item, big.NewInt(int64(v)))
            // fmt.Printf("done multiply\n")
        }
    case "+":
        if strings.Trim(val, " ") == "old" {
            // fmt.Printf("add\n")
            b.Add(item, item)
            // fmt.Printf("done add\n")
        } else {
            // fmt.Printf("strconv\n")
            v, _ := strconv.Atoi(val)
            // fmt.Printf("add\n")
            b.Add(item, big.NewInt(int64(v)))
            // fmt.Printf("done add\n")
        }
    }

    return b
}

func getBMonkeys(scanner *bufio.Scanner) []*BMonkey {
    monkeys := make([]*BMonkey, 0)
    for scanner.Scan() {
        line := scanner.Text()
        //monkey no
        n, _ := strconv.Atoi(string(line[7]))

        //starting items
        scanner.Scan()
        line = scanner.Text()
        line = line[17:]
        itemsStr := strings.Split(line, ",")
        items := make([]*big.Int, 0)
        for _, c := range itemsStr {
            // v, _ := strconv.Atoi(strings.Trim(c, " "))
            v, _ := new(big.Int).SetString(strings.Trim(c, " "), 10)
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

        monkey := NewBMonkey(n, operation, v, items, test, t, f)
        monkeys = append(monkeys, monkey)

        scanner.Scan()
    }
    return monkeys
}

