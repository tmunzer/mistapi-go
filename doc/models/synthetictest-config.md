
# Synthetictest Config

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Vlans` | [`[]models.SynthetictestProperties`](../../doc/models/synthetictest-properties.md) | Optional | - |
| `WanSpeedtest` | [`*models.SynthetictestConfigWanSpeedtest`](../../doc/models/synthetictest-config-wan-speedtest.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "vlans": [
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        "String7",
        "String8"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        "String7",
        "String8"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "wan_speedtest": {
    "enabled": false,
    "time_of_day": "time_of_day8",
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

