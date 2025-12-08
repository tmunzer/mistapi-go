
# Network Vpn Access Static Nat Property

## Structure

`NetworkVpnAccessStaticNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Static NAT destination IP Address. Must be an IP Address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "internal_ip": "192.168.70.3",
  "name": "pos_station-1"
}
```

