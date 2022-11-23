package hephaestus

import (
    "context"
    "testing"
)

func TestEnumRandomizeFunc(t *testing.T) {
    s := []string{"str", "str2", "str3"}
    v, err := EnumRandomizeFunc[string](s...)(context.TODO())
    if err != nil {
        t.Errorf("randomize func returned error: %s", err.Error())
    }

    for _, val := range s {
        if val == v {
            return
        }
    }

    if err != nil {
        t.Errorf("wrong value: %s, should be in slice %v", v, s)
    }
}

func TestEnumProbabilityRandomizeFunc(t *testing.T) {
    s := []EnumProbability[string]{
        {0.1, "str2"},
        {0.9, "str"},
        {0, "never"},
    }

    m := map[string]int{}
    n := 1000000

    for i := 0; i < n; i++ {
        v, err := EnumProbabilityRandomizeFunc[string](s...)(context.TODO())
        if err != nil {
            t.Errorf("randomize func returned error: %s", err.Error())
        }

        m[v]++
    }

    if len(m) > 2 {
        t.Errorf("randomized value with zero probability")
    }

    if float64(m["str"])/float64(n) < 0.89 {
        t.Errorf("invalid probability")
    }
}
