
# Service Max Latency

For SSR only, when `traffic_type`==`custom`, for uplink selection. 0-2147483647 or variable

## Class Name

`ServiceMaxLatency`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.ServiceMaxLatencyContainer.FromString(string mString) |
| `int` | models.ServiceMaxLatencyContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.ServiceMaxLatencyContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.ServiceMaxLatencyContainer.FromNumber(0)
```

