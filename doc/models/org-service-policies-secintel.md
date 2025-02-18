
# Org Service Policies Secintel

For SRX Only

*This model accepts additional fields of type interface{}.*

## Structure

`OrgServicePoliciesSecintel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | [`*models.SecintelProfileProfileActionEnum`](../../doc/models/secintel-profile-profile-action-enum.md) | Optional | enum: `default`, `standard`, `strict`<br>**Default**: `"default"` |
| `SecintelprofileId` | `*string` | Optional | org-level secintel Profile can be used, this takes precendence over 'profile' |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "profile": "default",
  "secintelprofile_id": "secintelprofile_id0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

