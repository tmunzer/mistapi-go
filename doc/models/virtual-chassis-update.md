
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
| `Op` | [`*models.VirtualChassisUpdateOpEnum`](../../doc/models/virtual-chassis-update-op-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "member": 118,
  "members": [
    {
      "mac": "mac2",
      "member": 176,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "linecard"
    },
    {
      "mac": "mac2",
      "member": 176,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "linecard"
    },
    {
      "mac": "mac2",
      "member": 176,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "linecard"
    }
  ],
  "new-member": 52,
  "op": "add"
}
```

