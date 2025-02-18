
# Wlan App Qos Apps Properties

*This model accepts additional fields of type interface{}.*

## Structure

`WlanAppQosAppsProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | `*int` | Optional | **Constraints**: `>= 0`, `<= 63` |
| `DstSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |
| `SrcSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dscp": 210,
  "dst_subnet": "dst_subnet0",
  "src_subnet": "src_subnet8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

