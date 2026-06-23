
# Radsec Idle Timeout

RadSec idle timeout in seconds. Default is 60

## Class Name

`RadsecIdleTimeout`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.RadsecIdleTimeoutContainer.FromNumber(int number) |
| `string` | models.RadsecIdleTimeoutContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.RadsecIdleTimeoutContainer.FromNumber(60)
```

## string

### Initialization Code

#### Example

```go
value := models.RadsecIdleTimeoutContainer.FromString("String0")
```

