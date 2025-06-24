
# Switch Oob Ip Config

Switch OOB IP Config:

- If HA configuration: key parameter will be nodeX (eg: node1)
- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | Optional, the network to be used for mgmt |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br><br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br><br>**Default**: `false` |
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
  "network": "network6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

