package hephaestus

import (
    "context"
    "math/rand"
)

func IntRandomizeFunc[T interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}](from, to T) RandomizeFunc[T] {
    if to < from {
        panic("to < from")
    }

    return func(ctx context.Context) (T, error) {
        return from + T(rand.Int())%(to-from), nil
    }
}

func FloatRandomizeFunc[T interface{ float32 | float64 }](from, to T) RandomizeFunc[T] {
    if to < from {
        panic("to < from")
    }

    k := to - from

    return func(ctx context.Context) (T, error) {
        return from + T(rand.Float64())*k, nil
    }
}
