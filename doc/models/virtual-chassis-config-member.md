
# Virtual Chassis Config Member

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisConfigMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional | - |
| `Mac` | `string` | Required | fpc0, same as the mac of device_id |
| `MemberId` | `*int` | Optional | to create a preprovisionned virtual chassis |
| `VcPorts` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `VcRole` | [`models.VirtualChassisConfigMemberVcRoleEnum`](../../doc/models/virtual-chassis-config-member-vc-role-enum.md) | Required | enum: `backup`, `linecard`, `master` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "locating": false,
  "mac": "mac2",
  "member_id": 216,
  "vc_ports": [
    "vc_ports2"
  ],
  "vc_role": "master",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

