
# Dhcp Client Option

*This model accepts additional fields of type interface{}.*

## Structure

`DhcpClientOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | - |
| `Data` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "code": "DHO_DHCP_MESSAGE_TYPE(53)",
  "data": "DHCPREQUEST",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

