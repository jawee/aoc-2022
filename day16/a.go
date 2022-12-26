package day16

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day16/testinput.txt")
    // file, err := os.Open(pwd + "/day16/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    m := make(map[string]Valve)
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        re := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnel[s]? lead[s]? to valve[s]? (.*)")
        match := re.FindStringSubmatch(scanner.Text())
        // for i, ma := range match {
        //     fmt.Printf("%d: %s\n", i, ma)
        // }

        name := match[1]
        fr, _ := strconv.Atoi(match[2])
        tunnels := strings.Split(strings.Replace(match[3]," ","", -1), ",")
        
        m[name] = Valve{fr, tunnels}
        fmt.Printf("%s name, fr %d, leads to %v. Len %d\n", name, fr, tunnels, len(tunnels))
    }
}

type Valve struct {
    flowRate int 
    edges []string
}

