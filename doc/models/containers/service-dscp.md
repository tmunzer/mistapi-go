
# Service Dscp

For SSR only, when `traffic_type`==`custom`. 0-63 or variable

## Class Name

`ServiceDscp`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.ServiceDscpContainer.FromString(string mString) |
| `int` | models.ServiceDscpContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.ServiceDscpContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.ServiceDscpContainer.FromNumber(0)
```

