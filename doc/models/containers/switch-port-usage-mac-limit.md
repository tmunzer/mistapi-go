
# Switch Port Usage Mac Limit

Only if `mode`!=`dynamic`, max number of MAC addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)

## Class Name

`SwitchPortUsageMacLimit`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SwitchPortUsageMacLimitContainer.FromNumber(int number) |
| `string` | models.SwitchPortUsageMacLimitContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMacLimitContainer.FromNumber(0)
```

## string

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMacLimitContainer.FromString("String0")
```

