
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
  "mac": "mac8",
  "member": 40,
  "vc_ports": [
    "vc_ports8",
    "vc_ports9",
    "vc_ports0"
  ],
  "vc_role": "master"
}
```

