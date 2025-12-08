
# Wxlan Tunnel Dmvpn

Dynamic Multipoint VPN configurations

## Structure

`WxlanTunnelDmvpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DMVPN is enabled<br><br>**Default**: `false` |
| `HoldingTime` | `*int` | Optional | Optional; the holding time for NHRP ‘registration requests’  and ‘resolution replies’ sent from the Mist AP (in seconds); default 600 |
| `HostRoutes` | `[]string` | Optional | Optional; list of IPv4 DMVPN peer host ip-addresses to which traffic is forwarded |

## Example (as JSON)

```json
{
  "enabled": false,
  "holding_time": 226,
  "host_routes": [
    "host_routes8",
    "host_routes9"
  ]
}
```

