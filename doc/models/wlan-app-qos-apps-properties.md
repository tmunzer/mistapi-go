
# Wlan App Qos Apps Properties

## Structure

`WlanAppQosAppsProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | [`*models.Dscp`](../../doc/models/containers/dscp.md) | Optional | DSCP value range between 0 and 63 |
| `DstSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |
| `SrcSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |

## Example (as JSON)

```json
{
  "dscp": "String3",
  "dst_subnet": "dst_subnet0",
  "src_subnet": "src_subnet8"
}
```

