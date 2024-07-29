
# Org Setting Auto Site Assignment

## Structure

`OrgSettingAutoSiteAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`models.Optional[[]models.OrgAutoRules]`](../../doc/models/org-auto-rules.md) | Optional | - |

## Example (as JSON)

```json
{
  "enable": false,
  "rules": [
    {
      "expression": "expression4",
      "match_device_type": "other",
      "match_model": "match_model0",
      "model": "model4",
      "prefix": "prefix6",
      "src": "name"
    },
    {
      "expression": "expression4",
      "match_device_type": "other",
      "match_model": "match_model0",
      "model": "model4",
      "prefix": "prefix6",
      "src": "name"
    }
  ]
}
```

