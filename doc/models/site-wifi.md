
# Site Wifi

Wi-Fi site settings

## Structure

`SiteWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiscoEnabled` | `*bool` | Optional | **Default**: `true` |
| `Disable11k` | `*bool` | Optional | whether to disable 11k<br>**Default**: `false` |
| `DisableRadiosWhenPowerConstrained` | `*bool` | Optional | **Default**: `false` |
| `EnableArpSpoofCheck` | `*bool` | Optional | when proxy_arp is enabled, check for arp spoofing.<br>**Default**: `false` |
| `EnableSharedRadioScanning` | `*bool` | Optional | **Default**: `true` |
| `Enabled` | `*bool` | Optional | enable WIFI feature (using SUB-MAN license)<br>**Default**: `true` |
| `LocateConnected` | `*bool` | Optional | whether to locate connected clients<br>**Default**: `true` |
| `LocateUnconnected` | `*bool` | Optional | whether to locate unconnected clients<br>**Default**: `false` |
| `MeshAllowDfs` | `*bool` | Optional | whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will distrupt the connection. If roaming is desired, keep it disabled.<br>**Default**: `false` |
| `MeshEnableCrm` | `*bool` | Optional | used to enable/disable CRM<br>**Default**: `false` |
| `MeshEnabled` | `*bool` | Optional | whether to enable Mesh feature for the site<br>**Default**: `false` |
| `MeshPsk` | `models.Optional[string]` | Optional | optional passphrase of mesh networking, default is generated randomly |
| `MeshSsid` | `models.Optional[string]` | Optional | optional ssid of mesh networking, default is based on site_id |
| `ProxyArp` | [`models.Optional[models.SiteWifiProxyArpEnum]`](../../doc/models/site-wifi-proxy-arp-enum.md) | Optional | enum: `default`, `disabled`, `enabled` |

## Example (as JSON)

```json
{
  "cisco_enabled": true,
  "disable_11k": false,
  "disable_radios_when_power_constrained": false,
  "enable_arp_spoof_check": false,
  "enable_shared_radio_scanning": true,
  "enabled": true,
  "locate_connected": true,
  "locate_unconnected": false,
  "mesh_allow_dfs": false,
  "mesh_enable_crm": false,
  "mesh_enabled": false
}
```

