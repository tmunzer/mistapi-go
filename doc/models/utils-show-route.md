
# Utils Show Route

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Rate at which output will refresh<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `Neighbor` | `*string` | Optional | IP of the neighbor |
| `Node` | [`*models.HaClusterNode`](../../doc/models/ha-cluster-node.md) | Optional | - |
| `Prefix` | `*string` | Optional | can be ip, ipv6 prefix |
| `Protocol` | [`*models.UtilsShowRouteProtocol2Enum`](../../doc/models/utils-show-route-protocol-2-enum.md) | Optional | (optional) Supported only for bgp<br><br>**Default**: `"bgp"` |
| `Route` | `*string` | Optional | if neighbor is specified, received / advertised; if not specified, both will be shown<br><br>* for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes<br>* for SRX and Switches, show route receive-protocol/advertise-protocol bgp 192.168.255.12 |
| `Vrf` | `*string` | Optional | VRF name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 0,
  "interval": 0,
  "neighbor": "192.168.4.1",
  "prefix": "192.168.0.5/30",
  "protocol": "bgp",
  "route": "advertised",
  "vrf": "default",
  "node": {
    "node": "node0",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

