
# Switch Mgmt Mxedge Proxy Port

Mist Edge port used to proxy the switch management traffic to the Mist Cloud. Value in range 1-65535

## Class Name

`SwitchMgmtMxedgeProxyPort`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.SwitchMgmtMxedgeProxyPortContainer.FromNumber(int number) |
| `string` | models.SwitchMgmtMxedgeProxyPortContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.SwitchMgmtMxedgeProxyPortContainer.FromNumber(2222)
```

## string

### Initialization Code

#### Example

```go
value := models.SwitchMgmtMxedgeProxyPortContainer.FromString("String0")
```

