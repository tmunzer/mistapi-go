
# Nac Rule

*This model accepts additional fields of type interface{}.*

## Structure

`NacRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.NacRuleActionEnum`](../../doc/models/nac-rule-action-enum.md) | Required | enum: `allow`, `block` |
| `ApplyTags` | `[]string` | Optional | all optional, this goes into Access-Accept |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | enabled or not<br>**Default**: `true` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Matching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `NotMatching` | [`*models.NacRuleMatching`](../../doc/models/nac-rule-matching.md) | Optional | - |
| `Order` | `*int` | Optional | the order of the rule, lower value implies higher priority<br>**Constraints**: `>= 0` |
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
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name0",
  "order": 1,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "created_time": 238.9,
  "matching": {
    "auth_type": "cert",
    "nactags": [
      "nactags6"
    ],
    "port_types": [
      "wired"
    ],
    "site_ids": [
      "00000738-0000-0000-0000-000000000000"
    ],
    "sitegroup_ids": [
      "00002472-0000-0000-0000-000000000000",
      "00002473-0000-0000-0000-000000000000",
      "00002474-0000-0000-0000-000000000000"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

