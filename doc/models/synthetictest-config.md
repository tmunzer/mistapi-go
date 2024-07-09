
# Synthetictest Config

## Structure

`SynthetictestConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Vlans` | [`[]models.SynthetictestProperties`](../../doc/models/synthetictest-properties.md) | Optional | - |

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
        186,
        187
      ]
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        186,
        187
      ]
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "vlan_ids": [
        186,
        187
      ]
    }
  ]
}
```
