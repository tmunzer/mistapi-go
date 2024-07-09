
# Virtual Chassis Config Member

## Structure

`VirtualChassisConfigMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | fpc0, same as the mac of device_id |
| `Member` | `*int` | Optional | to create a preprovisionned virtual chassis |
| `VcPorts` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `VcRole` | [`*models.VirtualChassisConfigMemberVcRoleEnum`](../../doc/models/virtual-chassis-config-member-vc-role-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "locating": false,
  "mac": "mac4",
  "member": 36,
  "vc_ports": [
    "vc_ports4",
    "vc_ports5"
  ],
  "vc_role": "master"
}
```

