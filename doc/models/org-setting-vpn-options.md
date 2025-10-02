
# Org Setting Vpn Options

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingVpnOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | **Constraints**: `>= 1`, `<= 2147483647` |
| `EnableIpv6` | `*bool` | Optional | **Default**: `false` |
| `StSubnet` | `*string` | Optional | requiring /12 or bigger to support 16 private IPs for 65535 gateways<br><br>**Default**: `"10.224.0.0/12"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enable_ipv6": false,
  "st_subnet": "10.224.0.0/12",
  "as_base": 80,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

