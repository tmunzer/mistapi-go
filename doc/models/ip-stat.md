
# Ip Stat

## Structure

`IpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | - |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Ip6` | `*string` | Optional | - |
| `Ips` | `map[string]string` | Optional | - |
| `Netmask` | `*string` | Optional | - |
| `Netmask6` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "gateway6": "fdad:b0bc:f29e::1",
  "ip": "10.3.3.1",
  "ip6": "fdad:b0bc:f29e::3d16",
  "netmask": "255.255.255.0",
  "netmask6": "/64",
  "dns": [
    "dns9"
  ],
  "dns_suffix": [
    "dns_suffix3"
  ],
  "gateway": "gateway4"
}
```

