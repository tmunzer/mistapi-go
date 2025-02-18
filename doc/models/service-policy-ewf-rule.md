
# Service Policy Ewf Rule

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicyEwfRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertOnly` | `*bool` | Optional | - |
| `BlockMessage` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Profile` | [`*models.ServicePolicyEwfRuleProfileEnum`](../../doc/models/service-policy-ewf-rule-profile-enum.md) | Optional | enum: `critical`, `standard`, `strict`<br>**Default**: `"strict"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "block_message": "Access to this URL Category has been blocked",
  "enabled": false,
  "profile": "strict",
  "alert_only": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

