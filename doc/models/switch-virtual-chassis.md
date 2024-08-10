
# Switch Virtual Chassis

required for preprovisioned Virtual Chassis

## Structure

`SwitchVirtualChassis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.SwitchVirtualChassisMember`](../../doc/models/switch-virtual-chassis-member.md) | Optional | list of Virtual Chassis members |
| `Preprovisioned` | `*bool` | Optional | to configure whether the VC is preprovisioned or nonprovisioned<br>**Default**: `false` |

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

