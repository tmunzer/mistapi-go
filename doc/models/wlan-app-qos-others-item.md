
# Wlan App Qos Others Item

## Structure

`WlanAppQosOthersItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | `*int` | Optional | - |
| `DstSubnet` | `*string` | Optional | - |
| `PortRanges` | `*string` | Optional | - |
| `Protocol` | `*string` | Optional | - |
| `SrcSubnet` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "dscp": 32,
  "dst_subnet": "10.2.0.0/16",
  "port_ranges": "80,1024-6553",
  "protocol": "udp",
  "src_subnet": "10.2.0.0/16"
}
```

