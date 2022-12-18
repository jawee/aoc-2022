package day13

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func B() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/day13/testinput.txt")
	file, err := os.Open(pwd + "/day13/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    dp1 := getObj("[[2]]")
    dp2 := getObj("[[6]]")
    objs := []Obj{dp1,dp2}

	for scanner.Scan() {
		l1 := scanner.Text()
		if l1 == "" {
			scanner.Scan()
			l1 = scanner.Text()
		}
		p1 := getObj(l1)
		objs = append(objs, p1)
	}

    sort.Slice(objs, func(i, j int) bool {
        res := getResult(objs[i],objs[j])
        if res == False {
            return false
        }
        return true
    })

    decoderKey := 1
    for i, o := range objs {
        if o.ToString() == "[[2]]" || o.ToString() == "[[6]]" {
            decoderKey *= (i+1)
        }
    }

	fmt.Printf("%d\n", decoderKey)
}
