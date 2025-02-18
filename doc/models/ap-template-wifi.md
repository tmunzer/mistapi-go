
# Ap Template Wifi

*This model accepts additional fields of type interface{}.*

## Structure

`ApTemplateWifi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CiscoEnabled` | `*bool` | Optional | - |
| `Disable11k` | `*bool` | Optional | **Default**: `false` |
| `DisableRadiosWhenPowerConstrained` | `*bool` | Optional | - |
| `EnableArpSpoof` | `*bool` | Optional | - |
| `EnableSharedRadioScanning` | `*bool` | Optional | **Default**: `false` |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `LocateConnected` | `*bool` | Optional | **Default**: `false` |
| `LocateUnconnected` | `*bool` | Optional | **Default**: `false` |
| `MeshAllowDfs` | `*bool` | Optional | **Default**: `false` |
| `MeshEnableCrm` | `*bool` | Optional | - |
| `MeshEnabled` | `*bool` | Optional | - |
| `ProxyArp` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_11k": false,
  "enable_shared_radio_scanning": false,
  "enabled": true,
  "locate_connected": false,
  "locate_unconnected": false,
  "mesh_allow_dfs": false,
  "proxy_arp": false,
  "cisco_enabled": false,
  "disable_radios_when_power_constrained": false,
  "enable_arp_spoof": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

