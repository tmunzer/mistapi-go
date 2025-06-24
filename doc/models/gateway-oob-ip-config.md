
# Gateway Oob Ip Config

Out-of-band (vme/em0/fxp0) IP config

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | If `type`==`static` |
| `Ip` | `*string` | Optional | If `type`==`static` |
| `Netmask` | `*string` | Optional | If `type`==`static` |
| `Node1` | [`*models.GatewayOobIpConfigNode1`](../../doc/models/gateway-oob-ip-config-node-1.md) | Optional | For HA Cluster, node1 can have different IP Config |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br><br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br><br>**Default**: `false` |
| `VlanId` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false,
  "gateway": "gateway0",
  "ip": "ip4",
  "netmask": "netmask0",
  "node1": {
    "gateway": "gateway2",
    "ip": "ip6",
    "netmask": "netmask2",
    "type": "dhcp",
    "use_mgmt_vrf": false,
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

