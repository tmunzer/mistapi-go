
# Switch Port Usage Mac Limit Overwrite

Max number of MAC addresses, default is 0 for unlimited, otherwise range is 1 to 16383 (upper bound constrained by platform)

## Class Name

`SwitchPortUsageMacLimitOverwrite`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SwitchPortUsageMacLimitOverwriteContainer.FromNumber(int number) |
| `string` | models.SwitchPortUsageMacLimitOverwriteContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMacLimitOverwriteContainer.FromNumber(0)
```

## string

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMacLimitOverwriteContainer.FromString("String0")
```

