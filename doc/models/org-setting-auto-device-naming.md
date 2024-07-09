
# Org Setting Auto Device Naming

## Structure

`OrgSettingAutoDeviceNaming`

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
      "match_device_type": "gateway",
      "match_model": "match_model0",
      "model": "model4",
      "prefix": "prefix6",
      "src": "model"
    },
    {
      "expression": "expression4",
      "match_device_type": "gateway",
      "match_model": "match_model0",
      "model": "model4",
      "prefix": "prefix6",
      "src": "model"
    }
  ]
}
```

