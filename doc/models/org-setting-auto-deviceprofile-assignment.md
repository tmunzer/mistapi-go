
# Org Setting Auto Deviceprofile Assignment

## Structure

`OrgSettingAutoDeviceprofileAssignment`

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
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "other",
      "src": "name"
    },
    {
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "other",
      "src": "name"
    },
    {
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "other",
      "src": "name"
    }
  ]
}
```

