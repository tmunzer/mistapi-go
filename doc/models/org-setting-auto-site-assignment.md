
# Org Setting Auto Site Assignment

## Structure

`OrgSettingAutoSiteAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`models.Optional[[]models.OrgSettingAutoAssignmentRule]`](../../doc/models/org-setting-auto-assignment-rule.md) | Optional | - |

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
      "match_device_type": "switch",
      "src": "name"
    },
    {
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "switch",
      "src": "name"
    }
  ]
}
```

