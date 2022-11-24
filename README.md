# hephaestus

![License](https://img.shields.io/github/license/squaddy-io/hephaestus) ![Tests status](https://img.shields.io/github/workflow/status/squaddy-io/hephaestus/test) ![go version](https://img.shields.io/github/go-mod/go-version/squaddy-io/hephaestus)

A package that provides cozy to use interface to randomize everything. Initially created to be used in NFT generation service.  
Feel free to use this in your projects.

## Usage
### Import
```go
import "github.com/squaddy-io/hephaestus"
```

### StructRandomizeFunc(m map[string]interface{})
Returns ready to use randomized object, where rules are put as map {`FieldName`: `RandomizeFunc`}.  
Panics on allocation if any field type mismatch occurred or any provided field is not found, and returns error if any provided `RandomizeFunc` returned an error.
#### Example
```go
fn := hephaestus.StructRandomizeFunc[Hero](map[string]interface{}{
    "Title": hephaestus.EnumRandomizeFunc[string](titles...),
    "Speed": hephaestus.IntRandomizeFunc[uint8](0, 120),
    "Rarity": hephaestus.EnumProbabilityRandomizeFunc[Rarity]([]hephaestus.EnumProbability[Rarity]{
        {0.9, CommonRarity}, {0.095, RareRarity}, {0.005, LegendaryRarity},
    }...),
})

...

nft, err := fn(ctx)
if err != nil {
    return nil, fmt.Errorf("failed to generate nft: %w", err)
}
```

#### Example with custom RandomizeFunc
```go
fn := hephaestus.StructRandomizeFunc[CustomStruct](map[string]interface{}{
    "ConstantField": func(_ context.Context) (string, error) { return "const" },
})
```

### IntRandomizeFunc(from, to T)
Returns any integer in `[from, to)`.  
Panics on allocation of `to < from` and never returns errors.
#### Example
```
color, _ := hephaestus.IntRandomizeFunc[int](0, 65535)(ctx)
```

### FloatRandomizeFunc(from, to T)
Returns any float number in `[from, to)`.  
Panics on allocation of `to < from` and never returns errors.
#### Example
```
charm, _ := hephaestus.IntRandomizeFunc[float32](0.5, 1)(ctx)
```

### EnumRandomizeFunc(s ...T)
Returns random value from slice `s`.  
Panics if `s` is empty and never returns errors.
#### Example
```
sex, _ := hephaestus.EnumRandomizeFunc[Sex](Male, Female)(ctx)
```

### EnumProbabilityRandomizeFunc(s ...EnumProbability)
Returns random value with set probability from slice `s`.  
Panics if `s` is empty or sum of probabilities is 0, and never returns errors.
#### Example
```
rarity, _ := hephaestus.EnumProbabilityRandomizeFunc[Rarity](
    []hephaestus.EnumProbability[Rarity]{
        {0.9, CommonRarity}, {0.095, RareRarity}, {0.005, LegendaryRarity},
    }...)(ctx)
```

### EnumCyclicRandomizeFunc(s ...T)
Iterates over slize `s` and returns values one by one.    
Panics if `s` is empty and never returns errors.
#### Example
```
team, _ := hephaestus.EnumCyclicRandomizeFunc[string]("red", "white", "black", "blue")(ctx)
```

### EnumLimitedRandomizeFunc(shuffle bool, s ...T)
Returns value (once for every element is slice `s`) from slice `s`. If parameter `shuffle` is true the function returns values in random order, if not - using order is slice.
Panics if `s` is empty and returns error if all elements were returned already returned.
#### Example
```
fn := hephaestus.EnumLimitedRandomizeFunc[string](true, freeLegendaryTitles...)(ctx)

...

title, err := fn(ctx)
if err != nil {
  return "", ErrAirdropFinished
}
```
