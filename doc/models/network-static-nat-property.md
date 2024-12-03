
# Network Static Nat Property

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkStaticNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `WanName` | `*string` | Optional | If not set, we configure the nat policies against all WAN ports for simplicity |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "internal_ip": "192.168.70.3",
  "name": "pos_station-1",
  "wan_name": "wan0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

