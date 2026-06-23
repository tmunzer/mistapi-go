
# Service Max Loss

For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-100 or variable

## Class Name

`ServiceMaxLoss`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.ServiceMaxLossContainer.FromString(string mString) |
| `int` | models.ServiceMaxLossContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.ServiceMaxLossContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.ServiceMaxLossContainer.FromNumber(0)
```

