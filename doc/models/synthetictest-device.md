
# Synthetictest Device

## Structure

`SynthetictestDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hostname` | `*string` | Optional | if `type`==`dns` |
| `Ip` | `*string` | Optional | if `type`==`arp` |
| `Password` | `*string` | Optional | if `type`==`radius` |
| `PortId` | `*string` | Optional | if `type`==`ssr` |
| `Type` | [`models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Required | enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `radius`, `speedtest` |
| `Url` | `*string` | Optional | if `type`==`curl` |
| `Username` | `*string` | Optional | if `type`==`radius` |
| `VlanId` | [`*models.SynthetictestDeviceVlanId`](../../doc/models/containers/synthetictest-device-vlan-id.md) | Optional | required for AP |

## Example (as JSON)

```json
{
  "hostname": "hostname2",
  "ip": "ip0",
  "password": "password0",
  "port_id": "port_id6",
  "type": "curl",
  "url": "url0"
}
```

