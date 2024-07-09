
# Utils Traceroute

## Structure

`UtilsTraceroute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | host name |
| `Network` | `*string` | Optional | for SSR, optional, the source to initiate traceroute from<br>**Default**: `"internal"` |
| `Port` | `*int` | Optional | when `protocol`==`udp`, the udp port to use<br>**Default**: `33434` |
| `Protocol` | [`*models.UtilsTracerouteProtocolEnum`](../../doc/models/utils-traceroute-protocol-enum.md) | Optional | **Default**: `"udp"` |
| `Timeout` | `*int` | Optional | maximum time in seconds to wait for the response<br>**Default**: `60` |
| `Vrf` | `*string` | Optional | for SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed |

## Example (as JSON)

```json
{
  "network": "internal",
  "port": 33434,
  "protocol": "udp",
  "timeout": 60,
  "host": "host4"
}
```

