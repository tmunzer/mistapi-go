
# Switch Matching Rule Oob Ip Config

Out-of-Band Management interface configuration

## Structure

`SwitchMatchingRuleOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | f supported on the platform. If enabled, DNS will be using this routing-instance, too |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": false,
  "use_mgmt_vrf_for_host_out": false
}
```

