
# Site Setting Tunterm Multicast Config

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.SiteSettingTuntermMulticastConfigMdns`](../../doc/models/site-setting-tunterm-multicast-config-mdns.md) | Optional | - |
| `MulticastAll` | `*bool` | Optional | **Default**: `false` |
| `Ssdp` | [`*models.SiteSettingTuntermMulticastConfigSsdp`](../../doc/models/site-setting-tunterm-multicast-config-ssdp.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "multicast_all": false,
  "mdns": {
    "enabled": false,
    "vlan_ids": [
      246,
      247
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ssdp": {
    "enabled": false,
    "vlan_ids": [
      236,
      237,
      238
    ],
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

