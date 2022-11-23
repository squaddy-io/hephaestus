package hephaestus

import (
    "context"
    "testing"
)

func Test_StructRandomizeFunc(t *testing.T) {
    type Role int
    const (
        WarriorRole Role = iota + 1
        MageRole
        DruidRole
    )

    type ConstType struct {
        String string
    }

    type SomeType struct {
        Tall   uint8
        Weight uint8
        Role   Role
        Const  ConstType
    }

    v, err := StructRandomizeFunc[SomeType](map[string]interface{}{
        "Tall": IntRandomizeFunc[uint8](120, 210),
        "Role": EnumRandomizeFunc[Role](WarriorRole, MageRole, DruidRole),
        "Const": func(_ context.Context) (ConstType, error) {
            return ConstType{String: "const"}, nil
        },
    })(context.TODO())

    if err != nil {
        t.Errorf("randomize func returned error: %s", err.Error())
    }

    if v.Tall == 0 {
        t.Errorf("Tall should be randomized")
    }
    if v.Weight != 0 {
        t.Errorf("Weight shouldn't be touched")
    }
    if v.Role == 0 {
        t.Errorf("Role should be randomized")
    }
    if v.Const.String != "const" {
        t.Errorf("Const.String should be constant")
    }
}
