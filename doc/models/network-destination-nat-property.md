
# Network Destination Nat Property

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkDestinationNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Port` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "internal_ip": "192.168.70.30",
  "name": "web server",
  "port": 443,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

