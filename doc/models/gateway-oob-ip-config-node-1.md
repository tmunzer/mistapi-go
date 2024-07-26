
# Gateway Oob Ip Config Node 1

for HA Cluster, node1 can have different IP Config

## Structure

`GatewayOobIpConfigNode1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | if `type`==`static` |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | used only if `subnet` is not specified in `networks` |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | **Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | if supported on the platform. If enabled, DNS will be using this routing-instance, too<br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | whether to use `mgmt_junos` for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br>**Default**: `false` |
| `VlanId` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false,
  "gateway": "gateway2",
  "ip": "ip6",
  "netmask": "netmask2"
}
```

