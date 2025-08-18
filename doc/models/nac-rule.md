
# Nac Rule

*This model accepts additional fields of type interface{}.*

## Structure

`NacRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.NacRuleActionEnum`](../../doc/models/nac-rule-action-enum.md) | Required | enum: `allow`, `block` |
| `ApplyTags` | `[]string` | Optional | All optional, this goes into Access-Accept |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | Enabled or not<br><br>**Default**: `true` |
| `GuestAuthState` | [`*models.NacRuleGuestAuthStateEnum`](../../doc/models/nac-rule-guest-auth-state-enum.md) | Optional | Guest portal authorization state. enum: `authorized`, `unknown` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Matching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `NotMatching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | - |
| `Order` | `*int` | Optional | Order of the rule, lower value implies higher priority<br><br>**Constraints**: `>= 0` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "allow",
  "apply_tags": [
    "c049dfcd-0c73-5014-1c64-062e9903f1e5"
  ],
  "enabled": true,
  "guest_auth_state": "authorized",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name0",
  "order": 1,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 238.9,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

