package hephaestus

import (
    "context"
    "fmt"
    "math/rand"
    "reflect"
    "time"
)

func init() {
    rand.Seed(time.Now().Unix())
}

type RandomizeFunc[T any] func(ctx context.Context) (T, error)

func StructRandomizeFunc[T any](m map[string]interface{}) RandomizeFunc[T] {
    t := reflect.TypeOf(*new(T))

    if t.Kind() != reflect.Struct {
        panic(fmt.Errorf("type should be a struct"))
    }

    for fieldName, randomizeFn := range m {
        ft, exist := t.FieldByName(fieldName)
        if !exist {
            panic(fmt.Errorf("%s field not found", fieldName))
        }

        fnv := reflect.ValueOf(randomizeFn)
        if fnv.Type().Kind() != reflect.Func {
            panic(fmt.Errorf("randomizer for %s field is not a func", fieldName))
        }

        if fnv.Type().NumIn() != 1 ||
            fnv.Type().In(0).Name() != "Context" ||
            fnv.Type().NumOut() != 2 ||
            fnv.Type().Out(0) != ft.Type ||
            fnv.Type().Out(1).Name() != "error" {
            panic(fmt.Errorf("randomizer for %s field is not a RandomizeFunc: should be func(ctx context.Context) (%s, error)", fieldName, ft.Type.Name()))
        }
    }

    return func(ctx context.Context) (ret T, err error) {
        v := reflect.ValueOf(new(T))
        for fieldName, fn := range m {
            if err = ctx.Err(); err != nil {
                return
            }

            out := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(context.Background())})
            if !out[1].IsNil() {
                err = fmt.Errorf("failed to randomize value for field %s: %w", fieldName, out[1].Interface().(error))
                return
            }

            reflect.Indirect(v).FieldByName(fieldName).Set(reflect.ValueOf(out[0].Interface()))
        }

        return reflect.Indirect(v).Interface().(T), nil
    }
}
