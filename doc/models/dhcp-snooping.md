
# Dhcp Snooping

*This model accepts additional fields of type interface{}.*

## Structure

`DhcpSnooping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | - |
| `EnableArpSpoofCheck` | `*bool` | Optional | Enable for dynamic ARP inspection check |
| `EnableIpSourceGuard` | `*bool` | Optional | Enable for check for forging source IP address |
| `Enabled` | `*bool` | Optional | - |
| `Networks` | `[]string` | Optional | If `all_networks`==`false`, list of network with DHCP snooping enabled |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "all_networks": false,
  "enable_arp_spoof_check": false,
  "enable_ip_source_guard": false,
  "enabled": false,
  "networks": [
    "networks8",
    "networks9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

