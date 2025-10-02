
# Dhcpd Config Fixed Binding

*This model accepts additional fields of type interface{}.*

## Structure

`DhcpdConfigFixedBinding`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Ip6` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip": "192.168.70.35",
  "ip6": "2607:f8b0:4005:808::2",
  "name": "name2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

