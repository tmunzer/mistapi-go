
# Switch Matching Rule Oob Ip Config

Out-of-Band Management interface configuration

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchMatchingRuleOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

