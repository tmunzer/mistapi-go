
# Wlan App Qos Others Item

## Structure

`WlanAppQosOthersItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | [`*models.Dscp`](../../doc/models/containers/dscp.md) | Optional | DSCP value range between 0 and 63 |
| `DstSubnet` | `*string` | Optional | - |
| `PortRanges` | `*string` | Optional | - |
| `Protocol` | `*string` | Optional | - |
| `SrcSubnet` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "dst_subnet": "10.2.0.0/16",
  "port_ranges": "80,1024-6553",
  "protocol": "udp",
  "src_subnet": "10.2.0.0/16",
  "dscp": "String3"
}
```

