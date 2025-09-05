
# Virtual Chassis Port

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.ConfigVcPortMember`](../../doc/models/config-vc-port-member.md) | Required | **Constraints**: *Unique Items Required* |
| `Op` | [`models.VirtualChassisPortOperationEnum`](../../doc/models/virtual-chassis-port-operation-enum.md) | Required | enum: `delete`, `set`<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "members": [
    {
      "member": 188.64,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "op": "delete",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

