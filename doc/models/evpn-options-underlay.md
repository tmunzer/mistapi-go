
# Evpn Options Underlay

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnOptionsUnderlay`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | Underlay BGP Base AS Number<br>**Default**: `65001`<br>**Constraints**: `>= 1`, `<= 65535` |
| `RoutedIdPrefix` | `*string` | Optional | - |
| `Subnet` | `*string` | Optional | Underlay subnet, by default, `10.255.240.0/20`, or `fd31:5700::/64` for ipv6 |
| `UseIpv6` | `*bool` | Optional | If v6 is desired for underlay<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "as_base": 65001,
  "routed_id_prefix": "/24",
  "subnet": "10.255.240.0/20",
  "use_ipv6": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

