
# Evpn Options Underlay

## Structure

`EvpnOptionsUnderlay`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | - |
| `RoutedIdPrefix` | `*string` | Optional | - |
| `Subnet` | `*string` | Optional | underlay subnet, by default, `10.255.240.0/20`, or `fd31:5700::/64` for ipv6 |
| `UseIpv6` | `*bool` | Optional | if v6 is desired for underlay<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "as_base": 65001,
  "routed_id_prefix": "/24",
  "subnet": "10.255.240.0/20",
  "use_ipv6": false
}
```

