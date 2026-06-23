
# Gateway Port Config Reth Idx

For SRX only and if HA Mode. `-1` means it will be managed by the device. Use `>= 0` values to manage it manually. Ensure no conflicting values are assigned across all ports.

## Class Name

`GatewayPortConfigRethIdx`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.GatewayPortConfigRethIdxContainer.FromNumber(int number) |
| `string` | models.GatewayPortConfigRethIdxContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.GatewayPortConfigRethIdxContainer.FromNumber(0)
```

## string

### Initialization Code

#### Example

```go
value := models.GatewayPortConfigRethIdxContainer.FromString("String0")
```

