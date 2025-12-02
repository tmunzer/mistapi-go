
# Mxedge Mgmt

## Structure

`MxedgeMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigAutoRevert` | `*bool` | Optional | **Default**: `false` |
| `FipsEnabled` | `*bool` | Optional | **Default**: `false` |
| `MistPassword` | `*string` | Optional | - |
| `OobIpType` | [`*models.MxedgeMgmtOobIpTypeEnum`](../../doc/models/mxedge-mgmt-oob-ip-type-enum.md) | Optional | enum: `dhcp`, `disabled`, `static`<br><br>**Default**: `"dhcp"` |
| `OobIpType6` | [`*models.MxedgeMgmtOobIpType6Enum`](../../doc/models/mxedge-mgmt-oob-ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"autoconf"` |
| `RootPassword` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "config_auto_revert": false,
  "fips_enabled": false,
  "mist_password": "MIST_PASSWORD",
  "oob_ip_type": "dhcp",
  "oob_ip_type6": "autoconf",
  "root_password": "ROOT_PASSWORD"
}
```

