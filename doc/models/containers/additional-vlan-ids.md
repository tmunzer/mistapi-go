
# Additional Vlan Ids

List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses

## Class Name

`AdditionalVlanIds`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.AdditionalVlanIdsContainer.FromString(string mString) |
| [`[]models.VlanIdWithVariable`](../../../doc/models/containers/vlan-id-with-variable.md) | models.AdditionalVlanIdsContainer.FromArrayOfVlanIdWithVariable5([]models.VlanIdWithVariable arrayOfVlanIdWithVariable5) |

## string

### Initialization Code

#### Example

```go
value := models.AdditionalVlanIdsContainer.FromString("String0")
```

## []models.VlanIdWithVariable

### Initialization Code

#### Example

```go
value := models.AdditionalVlanIdsContainer.FromArrayOfVlanIdWithVariable5([]models.VlanIdWithVariable{
    models.VlanIdWithVariableContainer.FromString("String8"),
})
```

