package hephaestus

import (
    "context"
    "strconv"
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

func TestEnumLimitedRandomizeFunc(t *testing.T) {
    var s []string
    for i := 0; i < 100; i++ {
        s = append(s, strconv.Itoa(i))
    }

    fn := EnumLimitedRandomizeFunc[string](false, s...)
    for i := 0; i < 100; i++ {
        v, err := fn(context.TODO())
        if err != nil {
            t.Errorf("pool shouldn't be drained")
        }

        if strconv.Itoa(i) != v {
            t.Fatalf("shouldn't be shuffled")
        }
    }

    _, err := fn(context.TODO())
    if err == nil {
        t.Errorf("pool should be drained")
    }
}

func TestEnumLimitedRandomizeFunc_Shuffle(t *testing.T) {
    var s []string
    for i := 0; i < 100; i++ {
        s = append(s, strconv.Itoa(i))
    }

    fn := EnumLimitedRandomizeFunc[string](true, s...)
    var shuffled bool
    for i := 0; i < 100; i++ {
        v, err := fn(context.TODO())
        if err != nil {
            t.Errorf("pool shouldn't be drained")
        }

        if strconv.Itoa(i) != v {
            shuffled = true
        }
    }

    _, err := fn(context.TODO())
    if err == nil {
        t.Errorf("pool should be drained")
    }

    if !shuffled {
        t.Errorf("should be shuffled")
    }
}

func TestEnumCyclicRandomizeFunc(t *testing.T) {
    fn := EnumCyclicRandomizeFunc[int](0, 1)

    for i := 0; i < 100; i++ {
        v, err := fn(context.TODO())
        if err != nil {
            t.Errorf("shouldn't return error: %s", err.Error())
        }

        if v != i%2 {
            t.Fatalf("expected %d got %d", i%2, v)
        }
    }
}
