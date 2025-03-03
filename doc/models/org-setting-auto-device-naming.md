
# Org Setting Auto Device Naming

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingAutoDeviceNaming`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enable` | `*bool` | Optional | - |
| `Rules` | [`models.Optional[[]models.OrgSettingAutoDeviceNamingRule]`](../../doc/models/org-setting-auto-device-naming-rule.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable": false,
  "rules": [
    {
      "expression": "expression4",
      "match_device": "ap",
      "prefix": "prefix6",
      "src": "lldp_port_desc",
      "suffix": "suffix2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "expression": "expression4",
      "match_device": "ap",
      "prefix": "prefix6",
      "src": "lldp_port_desc",
      "suffix": "suffix2",
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

