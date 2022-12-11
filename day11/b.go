package day11

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func B() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day11/testinput.txt")
    file, err := os.Open(pwd + "/day11/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()
        

    scanner := bufio.NewScanner(file)
    monkeys := getMonkeys(scanner)
    
    d := make([]int, 0)
    for _, m := range monkeys {
        d = append(d, m.test)
    }

    lcm := getLCMForList(d)
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
                item = calculateNewValue(m.op, m.v, item)
                // reduce item here in some way
                item := reduce(item, lcm)
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
    fmt.Printf("== After round 10000 ==\n")
    for _, m := range monkeys {
        // m.printMonkey()
        m.printCount()
        counts = append(counts, m.count)
    }

    sort.Ints(counts)
    
    fmt.Printf("%d\n", counts[len(counts)-1]*counts[len(counts)-2])

}

func getLCMForList(nums []int) int {
    res := 1
    for i := 0; i < len(nums); i++ {
        b := nums[i]
        res = getLCM(res,b)
    }
    return res
}

func getLCM(a, b int) int {
    x := a*b
    y := getGCD(a,b)
    return x/y
}

func getGCD(a,b int) int {
    if a == 0 {
        return b
    }
    return getGCD(b % a, a)
}

func reduce(item, i int) int {
    return item % i
}
