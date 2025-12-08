
# Stats Mxedge Oob Ip Stat

## Structure

`StatsMxedgeOobIpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | - |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Ip6` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | - |
| `Netmask6` | `*string` | Optional | - |
| `Type` | [`*models.MxedgeMgmtOobIpTypeEnum`](../../doc/models/mxedge-mgmt-oob-ip-type-enum.md) | Optional | enum: `dhcp`, `disabled`, `static`<br><br>**Default**: `"dhcp"` |
| `Type8` | [`*models.MxedgeMgmtOobIpType6Enum`](../../doc/models/mxedge-mgmt-oob-ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"autoconf"` |

## Example (as JSON)

```json
{
  "type": "dhcp",
  "type8": "autoconf",
  "dns": [
    "dns9",
    "dns8",
    "dns7"
  ],
  "gateway": "gateway0",
  "gateway6": "gateway66",
  "ip": "ip4",
  "ip6": "ip60"
}
```

