
# Switch Virtual Chassis

Required for preprovisioned Virtual Chassis

## Structure

`SwitchVirtualChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.SwitchVirtualChassisMember`](../../doc/models/switch-virtual-chassis-member.md) | Optional | List of Virtual Chassis members |
| `Preprovisioned` | `*bool` | Optional | To configure whether the VC is preprovisioned or nonprovisioned<br><br>**Default**: `false` |

## Example (as JSON)

```json
{
  "preprovisioned": false,
  "members": [
    {
      "mac": "mac2",
      "member_id": 58,
      "vc_role": "master"
    }
  ]
}
```

