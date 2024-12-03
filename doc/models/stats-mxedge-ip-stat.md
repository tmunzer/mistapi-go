
# Stats Mxedge Ip Stat

IP stats

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeIpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Ip6` | `*string` | Optional | - |
| `Ips` | `map[string]string` | Optional | Property key is the interface name. IPs for each net interface |
| `Macs` | `map[string]string` | Optional | Property key is the interface name. MAC for each net interface |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip": "192.168.1.244",
  "ip6": "fd4e:c615:b27d:5555::45",
  "ips": {
    "ens18": "92.168.1.244/24,fd4e:c615:b27d:5555::45/128,fd4e:c615:b27d:5555:20c:29ff:fe44:af25/64,fe80::104c:ffff:fee0:caf8/64"
  },
  "macs": {
    "ens18": "e4434b217044"
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

