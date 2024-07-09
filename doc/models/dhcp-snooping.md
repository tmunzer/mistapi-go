
# Dhcp Snooping

## Structure

`DhcpSnooping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | - |
| `EnableArpSpoofCheck` | `*bool` | Optional | Enable for dynamic ARP inspection check |
| `EnableIpSourceGuard` | `*bool` | Optional | Enable for check for forging source IP address |
| `Enabled` | `*bool` | Optional | - |
| `Networks` | `[]string` | Optional | if `all_networks`==`false`, list of network with DHCP snooping enabled |

## Example (as JSON)

```json
{
  "all_networks": false,
  "enable_arp_spoof_check": false,
  "enable_ip_source_guard": false,
  "enabled": false,
  "networks": [
    "networks8"
  ]
}
```

