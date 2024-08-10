
# Virtual Chassis Update

Virtual Chassis

## Structure

`VirtualChassisUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*int` | Optional | Only if `op`==`renumber` |
| `Members` | [`[]models.VirtualChassisMemberUpdate`](../../doc/models/virtual-chassis-member-update.md) | Optional | - |
| `NewMember` | `*int` | Optional | Only if `op`==`renumber` |
| `Op` | [`*models.VirtualChassisUpdateOpEnum`](../../doc/models/virtual-chassis-update-op-enum.md) | Optional | enum: `add`, `preprovision`, `remove`, `renumber` |

## Example (as JSON)

```json
{
  "member": 10,
  "members": [
    {
      "mac": "mac2",
      "member": 176,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "master"
    }
  ],
  "new-member": 200,
  "op": "add"
}
```

