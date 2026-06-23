
# Wlan Vlan Ids

WLAN VLAN pool IDs represented as either a comma-separated string or a list

## Class Name

`WlanVlanIds`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.WlanVlanIdsContainer.FromString(string mString) |
| [`[]models.VlanIdWithVariable`](../../../doc/models/containers/vlan-id-with-variable.md) | models.WlanVlanIdsContainer.FromArrayOfVlanIdWithVariable4([]models.VlanIdWithVariable arrayOfVlanIdWithVariable4) |

## string

### Initialization Code

#### Example

```go
value := models.WlanVlanIdsContainer.FromString("1,2")
```

## []models.VlanIdWithVariable

### Initialization Code

#### Example

```go
value := models.WlanVlanIdsContainer.FromArrayOfVlanIdWithVariable4([]models.VlanIdWithVariable{
    models.VlanIdWithVariableContainer.FromNumber(3),
    models.VlanIdWithVariableContainer.FromNumber(4),
    models.VlanIdWithVariableContainer.FromNumber(5),
})
```

