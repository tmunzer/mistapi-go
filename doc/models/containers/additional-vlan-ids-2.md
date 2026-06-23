
# Additional Vlan Ids 2

List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses

## Class Name

`AdditionalVlanIds2`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.AdditionalVlanIds2Container.FromString(string mString) |
| [`[]models.VlanIdWithVariable`](../../../doc/models/containers/vlan-id-with-variable.md) | models.AdditionalVlanIds2Container.FromArrayOfVlanIdWithVariable8([]models.VlanIdWithVariable arrayOfVlanIdWithVariable8) |

## string

### Initialization Code

#### Example

```go
value := models.AdditionalVlanIds2Container.FromString("String0")
```

## []models.VlanIdWithVariable

### Initialization Code

#### Example

```go
value := models.AdditionalVlanIds2Container.FromArrayOfVlanIdWithVariable8([]models.VlanIdWithVariable{
    models.VlanIdWithVariableContainer.FromString("String8"),
})
```

