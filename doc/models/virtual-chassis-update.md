
# Virtual Chassis Update

Virtual Chassis

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*int` | Optional | Only if `op`==`renumber` |
| `Members` | [`[]models.VirtualChassisMemberUpdate`](../../doc/models/virtual-chassis-member-update.md) | Optional | - |
| `NewMember` | `*int` | Optional | Only if `op`==`renumber` |
| `Op` | [`*models.VirtualChassisUpdateOpEnum`](../../doc/models/virtual-chassis-update-op-enum.md) | Optional | enum: `add`, `preprovision`, `remove`, `renumber` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "member": 10,
  "members": [
    {
      "mac": "mac2",
      "member": 176,
      "member_id": 58,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "master"
    }
  ],
  "new-member": 200,
  "op": "add",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

