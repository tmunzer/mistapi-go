
# Org Setting Auto Deviceprofile Assignment

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingAutoDeviceprofileAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`models.Optional[[]models.OrgSettingAutoAssignmentRule]`](../../doc/models/org-setting-auto-assignment-rule.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "src": "name",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "switch",
      "src": "name",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "create_new_site_if_needed": false,
      "expression": "expression4",
      "gatewaytemplate_id": "gatewaytemplate_id0",
      "match_country": "match_country8",
      "match_device_type": "switch",
      "src": "name",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

