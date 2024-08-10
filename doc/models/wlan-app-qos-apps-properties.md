
# Wlan App Qos Apps Properties

## Structure

`WlanAppQosAppsProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | `*int` | Optional | **Constraints**: `>= 0`, `<= 63` |
| `DstSubnet` | `*string` | Optional | subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |
| `SrcSubnet` | `*string` | Optional | subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |

## Example (as JSON)

```json
{
  "dscp": 210,
  "dst_subnet": "dst_subnet0",
  "src_subnet": "src_subnet8"
}
```

