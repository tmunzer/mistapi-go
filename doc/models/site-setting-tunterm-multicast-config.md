
# Site Setting Tunterm Multicast Config

## Structure

`SiteSettingTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.SiteSettingTuntermMulticastConfigMdns`](../../doc/models/site-setting-tunterm-multicast-config-mdns.md) | Optional | - |
| `MulticastAll` | `*bool` | Optional | **Default**: `false` |
| `Ssdp` | [`*models.SiteSettingTuntermMulticastConfigSsdp`](../../doc/models/site-setting-tunterm-multicast-config-ssdp.md) | Optional | - |

## Example (as JSON)

```json
{
  "multicast_all": false,
  "mdns": {
    "enabled": false,
    "vlan_ids": [
      246,
      247
    ]
  },
  "ssdp": {
    "enabled": false,
    "vlan_ids": [
      236,
      237,
      238
    ]
  }
}
```

