
# Stats Device Other Connected Device

*This model accepts additional fields of type interface{}.*

## Structure

`StatsDeviceOtherConnectedDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "020001abcdef",
  "name": "DNT-NTR-GWE",
  "port_id": "ge-0/0/1",
  "type": "gateway",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

