
# Virtual Chassis Member Update

## Structure

`VirtualChassisMemberUpdate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Required if `op`==`add` or `op`==`preprovision`. |
| `Member` | `*int` | Optional | Required if `op`==`remove` or `op`==`preprovision`. Optional if `op`==`add` |
| `VcPorts` | `[]string` | Optional | Required if `op`==`add` or `op`==`preprovision` |
| `VcRole` | [`*models.VirtualChassisMemberUpdateVcRoleEnum`](../../doc/models/virtual-chassis-member-update-vc-role-enum.md) | Optional | Required if `op`==`add` or `op`==`preprovision`. enum: `backup`, `linecard`, `master` |

## Example (as JSON)

```json
{
  "mac": "mac2",
  "member": 196,
  "vc_ports": [
    "vc_ports2",
    "vc_ports3"
  ],
  "vc_role": "backup"
}
```

