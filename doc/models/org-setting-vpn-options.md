
# Org Setting Vpn Options

## Structure

`OrgSettingVpnOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | **Constraints**: `>= 1`, `<= 4294967295` |
| `StSubnet` | `*string` | Optional | equiring /12 or bigger to support 16 private IPs for 65535 gateways<br>**Default**: `"10.224.0.0/12"` |

## Example (as JSON)

```json
{
  "st_subnet": "10.224.0.0/12",
  "as_base": 238
}
```

