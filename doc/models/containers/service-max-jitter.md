
# Service Max Jitter

For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable

## Class Name

`ServiceMaxJitter`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.ServiceMaxJitterContainer.FromString(string mString) |
| `int` | models.ServiceMaxJitterContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.ServiceMaxJitterContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.ServiceMaxJitterContainer.FromNumber(0)
```

