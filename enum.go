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
