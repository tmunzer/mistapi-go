
# Bgp Local As

Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. BGP AS, value in range 1-4294967295

## Class Name

`BgpLocalAs`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.BgpLocalAsContainer.FromString(string mString) |
| `int` | models.BgpLocalAsContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.BgpLocalAsContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.BgpLocalAsContainer.FromNumber(1)
```

