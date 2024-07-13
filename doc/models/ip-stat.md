
# Ip Stat

## Structure

`IpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpServer` | `models.Optional[string]` | Optional | - |
| `Dns` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Gateway` | `models.Optional[string]` | Optional | - |
| `Gateway6` | `models.Optional[string]` | Optional | - |
| `Ip` | `models.Optional[string]` | Optional | - |
| `Ip6` | `models.Optional[string]` | Optional | - |
| `Ips` | `map[string]string` | Optional | - |
| `Netmask` | `models.Optional[string]` | Optional | - |
| `Netmask6` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "dhcp_server": "192.168.95.1",
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

