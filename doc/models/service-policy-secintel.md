
# Service Policy Secintel

SRX only

## Structure

`ServicePolicySecintel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | [`*models.ServicePolicySecintelProfileEnum`](../../doc/models/service-policy-secintel-profile-enum.md) | Optional | enum: `default`, `standard`, `strict`<br><br>**Default**: `"default"` |
| `SecintelprofileId` | `*string` | Optional | org-level secintel Profile can be used, this takes precedence over 'profile' |

## Example (as JSON)

```json
{
  "enabled": false,
  "profile": "default",
  "secintelprofile_id": "secintelprofile_id4"
}
```

