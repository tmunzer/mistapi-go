
# Wlan Vlan Id with Variable

WLAN VLAN ID, either numeric, a variable string, or null

## Class Name

`WlanVlanIdWithVariable`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.WlanVlanIdWithVariableContainer.FromString(string mString) |
| `int` | models.WlanVlanIdWithVariableContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.WlanVlanIdWithVariableContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.WlanVlanIdWithVariableContainer.FromNumber(1)
```

