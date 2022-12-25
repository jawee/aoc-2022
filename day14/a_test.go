package day14

import (
	"fmt"
	"testing"
)

func TestDropSand(t *testing.T) {
    cases := []struct {
        input string
        x int
        n int
        expected []string
    }{
        {"1,2 -> 3,2", 2, 1, []string{"1,2","2,2","3,2","2,1"}},
        {"0,2 -> 4,2", 2, 2, []string{"0,2","1,2","2,2","3,2","4,2","2,1","1,1"}},
    }

    for _, c := range cases {
        fmt.Printf("#########################\n")
        fmt.Printf("Testcase drop sand %d times\n", c.n)
        fmt.Printf("#########################\n")

        m := make(map[string]int)
        getRocksForInput(c.input, m)

        for i := 0; i < c.n; i++ {
            _ = dropSand(c.x, m)
        }


        if len(m) != len(c.expected) {
            t.Fatalf("Expected length %d, got %d\n", len(c.expected), len(m))
        }
        for _, e := range c.expected {
            _, exists := m[e]
            if !exists {
                t.Fatalf("Expected obj at %s, got nil\n", e);
            }
        }

    }
}

func TestGetRocks(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {"1,1 -> 1,2", []string{"1,1","1,2"}},
        {"1,1 -> 1,3", []string{"1,1","1,2","1,3"}},
        {"1,1 -> 1,2 -> 2,2", []string{"1,1","1,2","2,2"}},
        {"1,1 -> 1,3 -> 2,3", []string{"1,1","1,2","1,3","2,3"}},
        {"3,3 -> 3,1 -> 1,1", []string{"3,3","3,2","3,1","2,1","1,1"}},
    }

    for _, c := range cases {
        fmt.Printf("#########################\n")
        fmt.Printf("Testcase %s\n", c.input)
        fmt.Printf("#########################\n")
        m := make(map[string]int)
        getRocksForInput(c.input, m)

        if len(m) != len(c.expected) {
            t.Fatalf("Expected length %d, got %d\n", len(c.expected), len(m))
        }
        for _, e := range c.expected {
            _, exists := m[e]
            if !exists {
                t.Fatalf("Expected rock at %s, got nil\n", e);
            }
        }
    }
}
