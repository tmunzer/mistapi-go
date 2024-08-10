
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
| `VcRole` | [`*models.VirtualChassisConfigMemberVcRoleEnum`](../../doc/models/virtual-chassis-config-member-vc-role-enum.md) | Optional | enum: `backup`, `linecard`, `master` |

## Example (as JSON)

```json
{
  "locating": false,
  "mac": "mac2",
  "member": 194,
  "vc_ports": [
    "vc_ports2"
  ],
  "vc_role": "master"
}
```

