
# Mxedge Mgmt

## Structure

`MxedgeMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FipsEnabled` | `*bool` | Optional | **Default**: `false` |
| `MistPassword` | `*string` | Optional | - |
| `OobIpType` | [`*models.MxedgeMgmtOobIpTypeEnum`](../../doc/models/mxedge-mgmt-oob-ip-type-enum.md) | Optional | **Default**: `"dhcp"` |
| `OobIpType6` | [`*models.MxedgeMgmtOobIpType6Enum`](../../doc/models/mxedge-mgmt-oob-ip-type-6-enum.md) | Optional | **Default**: `"autoconf"` |
| `RootPassword` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "fips_enabled": false,
  "mist_password": "MIST_PASSWORD",
  "oob_ip_type": "dhcp",
  "oob_ip_type6": "autoconf",
  "root_password": "ROOT_PASSWORD"
}
```

