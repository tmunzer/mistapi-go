
# Bgp As

BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` )

## Class Name

`BgpAs`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.BgpAsContainer.FromString(string mString) |
| `int` | models.BgpAsContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.BgpAsContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.BgpAsContainer.FromNumber(1)
```

