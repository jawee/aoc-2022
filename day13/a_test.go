package day13

import "testing"

func TestGetResult(t *testing.T) {
    cases := []struct {str1 string; str2 string; expected Result}{
        // {"[1,1,3,1,1]", "[1,1,5,1,1]", True},
        // {"[[1]]", "[[2]]", True},
        // {"[[2]]", "[[1]]", False},
        // {"[1]", "[[2,3]]", True},
        // {"[9]", "[[8,7]]", False},
        // {"[]", "[1]",True},
        // {"[1]", "[]",False},
        // {"[[]]", "[]",False},
        // {"[[[]]]", "[]",False},
        // {"[[[[]]]]", "[]",False},
        // {"[[[[]]]]", "[[]]",False},
        // {"[7,7,7,7]", "[7,7,7]",False},
        // {"[7,7,7,7]", "[7,7,7,7]",True},
        // {"[7,7,7,7]", "[7,7,7,7,7]",True},
        // {"[10]", "[11]", True},
        // {"[10]", "[9]", False},
        // {"[[[[],[10],[5,6,2],6],[2,[6,3],7,5,2],[6,10,5,[],6],[[5,9],5]],[],[[10],[[8,9,1,7],[0,8,10,10]],8],[2,[[0,5],[0,6],[],0,[7,8,4]],[[4,1,7,6],4,8,[],[5,7,7,0]],10],[[[],0,4,3],5]]", "[[5,10,[[4,5],8,[0,7,5]],[6,[5,6,4,0,7],1]]]", False},
        // {"[[[]]]", "[[]]", False},
        // { "[]", "[1]", True },
        // { "[[]]", "[1]", True },
        // { "[[[]]]", "[[1]]", True },
        // {"[[[[]]]]", "[[3],[8,10]]",True},
        // {"[[8,[[7]]]]", "[[[[8],2]]]",True},
        // {"[[8,[[7]]]]", "[[[[8]]]]", False},
        // {"[[8,[[7]]]]", "[[[[8],[3]]]]",False},
        // {"[[1,2],4]", "[[1],5,5]", False},
        // {"[[1,2],4]", "[[[3]],5,5]", True},
        // {"[1,2,3,[1,2,3],4,1]", "[1,2,3,[1,2,3],4,0]", False},
        {"[[8,[[7]]]]", "[[[[8]]]]", False},
        {"[[8,[[7]]]]", "[[[[8],2]]]",True},
    }

    for _, c := range cases {
        p1 := getObj(c.str1)
        p2 := getObj(c.str2)

        // pair := Pair{p1,p2,c.expected}

        res := getResult(p1,p2)

        if res != c.expected {
            t.Fatalf("%s vs %s. Expected %s, got %s\n", c.str1, c.str2, resultToString(c.expected), resultToString(res))
        }
    }
}

func TestPrint(t *testing.T) {
    cases := []struct {str string}{
        {"[]"},
        {"[[]]"},
        {"[1]"},
        {"[1,2]"},
        {"[1,2,2]"},
        {"[[1],2,2]"},
    }

    for _, c := range cases {
        obj := getObj(c.str)
        res := obj.ToString()

        if res != c.str {
            t.Fatalf("Expected %s, got %s\n", c.str, res)
        }
    }
}
func resultToString(a Result) string {
    switch a {
    case True:
        return "True"
    case False:
        return "False"
    case Continue:
        return "Continue"
    }

    return "unknown result"
}
