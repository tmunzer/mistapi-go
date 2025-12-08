
# Site Wifi

Wi-Fi site settings

## Structure

`SiteWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiscoEnabled` | `*bool` | Optional | **Default**: `true` |
| `Disable11k` | `*bool` | Optional | Whether to disable 11k<br><br>**Default**: `false` |
| `DisableRadiosWhenPowerConstrained` | `*bool` | Optional | **Default**: `false` |
| `EnableArpSpoofCheck` | `*bool` | Optional | When proxy_arp is enabled, check for arp spoofing.<br><br>**Default**: `false` |
| `EnableSharedRadioScanning` | `*bool` | Optional | **Default**: `true` |
| `Enabled` | `*bool` | Optional | Enable Wi-Fi feature (using SUB-MAN license)<br><br>**Default**: `true` |
| `LocateConnected` | `*bool` | Optional | Whether to locate connected clients<br><br>**Default**: `true` |
| `LocateUnconnected` | `*bool` | Optional | Whether to locate unconnected clients<br><br>**Default**: `false` |
| `MeshAllowDfs` | `*bool` | Optional | Whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will disrupt the connection. If roaming is desired, keep it disabled.<br><br>**Default**: `false` |
| `MeshEnableCrm` | `*bool` | Optional | Used to enable/disable CRM<br><br>**Default**: `false` |
| `MeshEnabled` | `*bool` | Optional | Whether to enable Mesh feature for the site<br><br>**Default**: `false` |
| `MeshPsk` | `models.Optional[string]` | Optional | Optional passphrase of mesh networking, default is generated randomly |
| `MeshSsid` | `models.Optional[string]` | Optional | Optional ssid of mesh networking, default is based on site_id |
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

