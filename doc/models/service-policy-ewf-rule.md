
# Service Policy Ewf Rule

## Structure

`ServicePolicyEwfRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertOnly` | `*bool` | Optional | - |
| `BlockMessage` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | [`*models.ServicePolicyEwfRuleProfileEnum`](../../doc/models/service-policy-ewf-rule-profile-enum.md) | Optional | enum: `critical`, `standard`, `strict`<br><br>**Default**: `"strict"` |

## Example (as JSON)

```json
{
  "block_message": "Access to this URL Category has been blocked",
  "enabled": false,
  "profile": "strict",
  "alert_only": false
}
```

