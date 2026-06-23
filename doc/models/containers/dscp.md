
# Dscp

DSCP value range between 0 and 63

## Class Name

`Dscp`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.DscpContainer.FromString(string mString) |
| `int` | models.DscpContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.DscpContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.DscpContainer.FromNumber(0)
```

