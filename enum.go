package hephaestus

import (
    "context"
    "fmt"
    "math/rand"
)

func EnumRandomizeFunc[T any](s ...T) RandomizeFunc[T] {
    if len(s) == 0 {
        panic(fmt.Errorf("empty slice"))
    }

    return func(ctx context.Context) (T, error) {
        return s[rand.Int()%len(s)], nil
    }
}

type EnumProbability[T any] struct {
    Probability float64
    Value       T
}

func EnumProbabilityRandomizeFunc[T any](s ...EnumProbability[T]) RandomizeFunc[T] {
    if len(s) == 0 {
        panic(fmt.Errorf("empty slice"))
    }

    var k float64
    for _, v := range s {
        k += v.Probability
    }
    if k == 0 {
        panic(fmt.Errorf("invalid probabilities slice"))
    }

    return func(ctx context.Context) (ret T, err error) {
        r := rand.Float64() * k

        for _, v := range s {
            if r < v.Probability {
                return v.Value, nil
            }
            r -= v.Probability
        }

        return s[len(s)-1].Value, nil
    }
}

func EnumCyclicRandomizeFunc[T any](s ...T) RandomizeFunc[T] {
    if len(s) == 0 {
        panic(fmt.Errorf("empty slice"))
    }

    var i int

    return func(ctx context.Context) (out T, err error) {
        out = s[i]
        i++

        if len(s) == i {
            i = 0
        }

        return
    }
}

func EnumLimitedRandomizeFunc[T any](shuffle bool, s ...T) RandomizeFunc[T] {
    if len(s) == 0 {
        panic(fmt.Errorf("pool is empty"))
    }

    if shuffle {
        rand.Shuffle(len(s), func(i, j int) {
            s[i], s[j] = s[j], s[i]
        })
    }

    return func(ctx context.Context) (T, error) {
        var out T

        if len(s) == 0 {
            return out, fmt.Errorf("pool is empty")
        }

        out, s = s[0], s[1:]
        return out, nil
    }
}
