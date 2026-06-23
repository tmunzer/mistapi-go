
# Vlan Id with Variable

VLAN ID, either numeric or expressed as a template variable string

## Class Name

`VlanIdWithVariable`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.VlanIdWithVariableContainer.FromString(string mString) |
| `int` | models.VlanIdWithVariableContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.VlanIdWithVariableContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.VlanIdWithVariableContainer.FromNumber(1)
```

