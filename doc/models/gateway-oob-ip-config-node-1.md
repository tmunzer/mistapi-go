
# Gateway Oob Ip Config Node 1

For HA Cluster, node1 can have different IP Config

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayOobIpConfigNode1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | If `type`==`static` |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks` |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | Whether to use `mgmt_junos` for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br>**Default**: `false` |
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

