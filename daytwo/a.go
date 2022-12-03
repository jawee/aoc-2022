package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/daytwo/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalScore := 0
    for scanner.Scan() {
        game := strings.Split(scanner.Text(), " ")
        opponent := game[0]
        player := game[1]
        totalScore = totalScore + getScoreForChoice(player) + getScoreForGame(opponent, player)
    }
    
    fmt.Println(totalScore)
}

/* 
Rock = A X = 1
Paper = B Y = 2 
Scissor = C Z = 3

win = 6
draw = 3
loss = 0
*/

func getMove(a string) string {
    move := ""
    switch a {
    case "A":
        move = "ROCK"
    case "X":
        move = "ROCK"
    case "B": 
        move = "PAPER"
    case "Y":
        move = "PAPER"
    case "C":
        move = "SCISSOR"
    case "Z":
        move = "SCISSOR"
    }

    return move
}

func getScoreForGame(a string, b string) int {
    a = getMove(a)
    b = getMove(b)
    fmt.Println(a + " " + b)
    if a == b {
        fmt.Println("draw")
        return 3
    }

    if b == "ROCK" && a == "SCISSOR" {
        fmt.Println("win")
        return 6
    }

    if b == "PAPER" && a == "ROCK" {
        fmt.Println("win")
        return 6
    }

    if b == "SCISSOR" && a == "PAPER" {
        fmt.Println("win")
        return 6
    }

    fmt.Println("loss")
    return 0
}

func getScoreForChoice(choice string) int {
    switch choice {
        case "C":
            return 3
        case "B":
            return 2
        case "A":
            return 1
        case "X":
            return 1
        case "Y":
            return 2
        case "Z":
            return 3
    }

    return 0
}
