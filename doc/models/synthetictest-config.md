
# Synthetictest Config

## Structure

`SynthetictestConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Vlans` | [`[]models.SynthetictestProperties`](../../doc/models/synthetictest-properties.md) | Optional | - |
| `WanSpeedtest` | [`*models.SynthetictestConfigWanSpeedtest`](../../doc/models/synthetictest-config-wan-speedtest.md) | Optional | - |

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
      ]
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        "String7",
        "String8"
      ]
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        "String7",
        "String8"
      ]
    }
  ],
  "wan_speedtest": {
    "enabled": false,
    "time_od_fay": "time_od_fay2"
  }
}
```

