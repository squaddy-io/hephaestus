package hephaestus

import (
    "context"
    "testing"
)

func TestIntRandomizeFunc(t *testing.T) {
    v, err := IntRandomizeFunc[uint8](120, 210)(context.TODO())
    if err != nil {
        t.Errorf("randomize func returned error: %s", err.Error())
    }

    if v < 120 || v > 209 {
        t.Errorf("val should be between 120 and 209, got %d", v)
    }
}

func TestFloatRandomizeFunc(t *testing.T) {
    v, err := FloatRandomizeFunc[float32](0.5, 2.5)(context.TODO())
    if err != nil {
        t.Errorf("randomize func returned error: %s", err.Error())
    }

    if v < 0.5 || v >= 2.5 {
        t.Errorf("val should be between 0.5 and 2.5, got %f", v)
    }
}
