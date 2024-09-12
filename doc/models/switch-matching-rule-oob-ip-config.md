
# Switch Matching Rule Oob Ip Config

Out-of-Band Management interface configuration

## Structure

`SwitchMatchingRuleOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*interface{}` | Optional | f supported on the platform. If enabled, DNS will be using this routing-instance, too |
| `UseMgmtVrfForHostOut` | `*interface{}` | Optional | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired, |

## Example (as JSON)

```json
{
  "type": "static",
  "use_mgmt_vrf": {
    "key1": "val1",
    "key2": "val2"
  },
  "use_mgmt_vrf_for_host_out": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

