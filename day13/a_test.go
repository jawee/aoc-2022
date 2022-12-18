package day13

import "testing"

func TestGetResult(t *testing.T) {
    cases := []struct {str1 string; str2 string; expected int}{
        {"[1,1,3,1,1]", "[1,1,5,1,1]", 1},
        {"[[1]]", "[[2]]", 2},
        {"[[2]]", "[[1]]", 0},
        {"[1]", "[[2,3]]", 3},
        {"[1]", "[[1,2]]", 4},
        {"[9]", "[[8,7]]", 0},
        {"[]", "[1]",5},
        {"[1]", "[]",0},
        {"[[]]", "[]",0},
        {"[[[]]]", "[]",0},
        {"[[[[]]]]", "[]",0},
        {"[[[[]]]]", "[[]]",0},
        {"[7,7,7,7]", "[7,7,7]",0},
        {"[7,7,7,7]", "[7,7,7,7]",6},
        {"[7,7,7,7]", "[7,7,7,7,7]",7},
        {"[10]", "[11}", 8},
    }

    for _, c := range cases {
        p1 := getObj(c.str1)
        p2 := getObj(c.str2)

        pair := Pair{p1,p2,c.expected}

        res := getResult(pair)

        if res != c.expected {
            t.Fatalf("%s vs %s. Expected %d, got %d\n", c.str1, c.str2, c.expected, res)
        }
    }
}
