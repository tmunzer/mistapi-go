
# Switch Virtual Chassis

Required for preprovisioned Virtual Chassis

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchVirtualChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.SwitchVirtualChassisMember`](../../doc/models/switch-virtual-chassis-member.md) | Optional | List of Virtual Chassis members |
| `Preprovisioned` | `*bool` | Optional | To configure whether the VC is preprovisioned or nonprovisioned<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "preprovisioned": false,
  "members": [
    {
      "mac": "mac2",
      "member_id": 58,
      "vc_role": "master",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

