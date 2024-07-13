
# Gateway Oob Ip Config

out-of-band (vme/em0/fxp0) IP config

## Structure

`GatewayOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | optional, the network to be used for mgmt |
| `Node1` | [`*models.GatewayOobIpConfigNode1`](../../doc/models/gateway-oob-ip-config-node-1.md) | Optional | for HA Cluster, node1 can have different IP Config |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | **Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | f supported on the platform. If enabled, DNS will be using this routing-instance, too<br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired,<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false,
  "ip": "ip8",
  "netmask": "netmask4",
  "network": "network0",
  "node1": {
    "ip": "ip6",
    "netmask": "netmask2",
    "network": "network8",
    "type": "static",
    "use_mgmt_vrf": false
  }
}
```

