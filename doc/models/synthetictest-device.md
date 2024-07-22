
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
| `Type` | [`models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Required | - |
| `Url` | `*string` | Optional | if `type`==`curl` |
| `Username` | `*string` | Optional | if `type`==`radius` |
| `VlanId` | [`*models.SynthetictestDeviceVlanId`](../../doc/models/containers/synthetictest-device-vlan-id.md) | Optional | This is a container for one-of cases. |

## Example (as JSON)

```json
{
  "hostname": "hostname8",
  "ip": "ip0",
  "password": "password0",
  "port_id": "port_id6",
  "type": "curl",
  "url": "url0"
}
```

