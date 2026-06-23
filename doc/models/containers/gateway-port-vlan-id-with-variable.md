
# Gateway Port Vlan Id with Variable

If WAN interface is on a VLAN. Can be the VLAN ID (i.e. "10") or a Variable (i.e. "{{myvar}}")

## Class Name

`GatewayPortVlanIdWithVariable`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.GatewayPortVlanIdWithVariableContainer.FromString(string mString) |
| `int` | models.GatewayPortVlanIdWithVariableContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.GatewayPortVlanIdWithVariableContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.GatewayPortVlanIdWithVariableContainer.FromNumber(1)
```

