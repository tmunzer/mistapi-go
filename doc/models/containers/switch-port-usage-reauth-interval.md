
# Switch Port Usage Reauth Interval

Only if `mode`!=`dynamic` and `port_auth`=`dot1x` reauthentication interval range (min: 10, max: 65535, default: 3600). Set to 0 to disable reauthentication (no-reauthentication).

## Class Name

`SwitchPortUsageReauthInterval`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SwitchPortUsageReauthIntervalContainer.FromNumber(int number) |
| `string` | models.SwitchPortUsageReauthIntervalContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageReauthIntervalContainer.FromNumber(3600)
```

## string

### Initialization Code

#### Example

```go
value := models.SwitchPortUsageReauthIntervalContainer.FromString("String0")
```

