
# Wxlan Tunnel Dmvpn

Dynamic Multipoint VPN configurations

## Structure

`WxlanTunnelDmvpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether DMVPN is enabled<br>**Default**: `false` |
| `HoldingTime` | `*int` | Optional | optional; the holding time for NHRP ‘registration requests’ and ‘resolution replies’ sent from the Mist AP (in seconds); default 600 |
| `HostRoutes` | `[]string` | Optional | optional; list of IPv4 DMVPN peer host ip-addresses to which traffic is forwarded |

## Example (as JSON)

```json
{
  "enabled": false,
  "holding_time": 156,
  "host_routes": [
    "host_routes2",
    "host_routes3"
  ]
}
```

