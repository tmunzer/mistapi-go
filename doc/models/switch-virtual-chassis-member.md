
# Switch Virtual Chassis Member

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchVirtualChassisMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | fpc0, same as the mac of device_id |
| `MemberId` | `*int` | Optional | - |
| `VcRole` | [`*models.SwitchVirtualChassisMemberVcRoleEnum`](../../doc/models/switch-virtual-chassis-member-vc-role-enum.md) | Optional | Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config. enum: `backup`, `linecard`, `master` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "aff827549235",
  "member_id": 104,
  "vc_role": "backup",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

