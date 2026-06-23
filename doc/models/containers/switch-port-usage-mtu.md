
# Switch Port Usage Mtu

Only if `mode`!=`dynamic` media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation. The default value is 1514.

## Class Name

`SwitchPortUsageMtu`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SwitchPortUsageMtuContainer.FromNumber(int number) |
| `string` | models.SwitchPortUsageMtuContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMtuContainer.FromNumber(256)
```

## string

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageMtuContainer.FromString("String0")
```

