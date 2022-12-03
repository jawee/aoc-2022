package dayone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/dayone/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()


    scanner := bufio.NewScanner(file)
    currSum := 0
    var list []int
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        if scanner.Text() == "" {
            list = append(list, currSum)
            currSum = 0
        }
        v, _ := strconv.Atoi(scanner.Text())
        currSum = currSum + v
    }
    list = append(list, currSum)

    sort.Slice(list, func(i, j int) bool {
        return list[i] < list[j]
    })

    
    list = list[len(list)-3:]
 
    sum := 0
    for _, v := range list {
        sum = sum + v
    }

    fmt.Println(sum)
}
