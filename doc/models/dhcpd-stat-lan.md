
# Dhcpd Stat Lan

*This model accepts additional fields of type interface{}.*

## Structure

`DhcpdStatLan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NumIps` | `*int` | Optional | - |
| `NumLeased` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "num_ips": 100,
  "num_leased": 20,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

