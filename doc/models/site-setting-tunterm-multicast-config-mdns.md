
# Site Setting Tunterm Multicast Config Mdns

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingTuntermMulticastConfigMdns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `VlanIds` | `[]int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "vlan_ids": [
    2,
    3,
    5
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

