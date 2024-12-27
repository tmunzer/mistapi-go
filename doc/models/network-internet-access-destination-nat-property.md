
# Network Internet Access Destination Nat Property

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkInternetAccessDestinationNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Destination NAT destination IP Address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | - |
| `Port` | `*string` | Optional | The Destination NAT destination IP Address. Must be a Port (i.e. "443") or a Variable (i.e. "{{myvar}}") |
| `WanName` | `*string` | Optional | SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "internal_ip": "192.168.70.30",
  "name": "web server",
  "port": "443",
  "wan_name": "wan0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

