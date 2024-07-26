
# Gateway Oob Ip Config

out-of-band (vme/em0/fxp0) IP config

## Structure

`GatewayOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | if `type`==`static` |
| `Ip` | `*string` | Optional | if `type`==`static` |
| `Netmask` | `*string` | Optional | if `type`==`static` |
| `Node1` | [`*models.GatewayOobIpConfigNode1`](../../doc/models/gateway-oob-ip-config-node-1.md) | Optional | for HA Cluster, node1 can have different IP Config |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | **Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | if supported on the platform. If enabled, DNS will be using this routing-instance, too<br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired,<br>**Default**: `false` |
| `VlanId` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false,
  "gateway": "gateway4",
  "ip": "ip8",
  "netmask": "netmask4",
  "node1": {
    "gateway": "gateway2",
    "ip": "ip6",
    "netmask": "netmask2",
    "type": "static",
    "use_mgmt_vrf": false
  }
}
```

