
# Wxlan Tunnel Dmvpn

Dynamic Multipoint VPN configurations

*This model accepts additional fields of type interface{}.*

## Structure

`WxlanTunnelDmvpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DMVPN is enabled<br>**Default**: `false` |
| `HoldingTime` | `*int` | Optional | Optional; the holding time for NHRP ‘registration requests’  and ‘resolution replies’ sent from the Mist AP (in seconds); default 600 |
| `HostRoutes` | `[]string` | Optional | Optional; list of IPv4 DMVPN peer host ip-addresses to which traffic is forwarded |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "holding_time": 226,
  "host_routes": [
    "host_routes8",
    "host_routes9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

