
# Mxedge Stats Ip Stat

OOBM IP stats

## Structure

`MxedgeStatsIpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Ips` | `map[string]string` | Optional | Property key is the interface name. IPs for each net interface |
| `Macs` | `map[string]string` | Optional | Property key is the interface name. MAC for each net interface |

## Example (as JSON)

```json
{
  "ip": "192.168.1.244",
  "ips": {
    "ens18": "192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"
  },
  "macs": {
    "ens18": "e4434b217044"
  }
}
```

