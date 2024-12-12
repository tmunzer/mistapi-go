
# Stats Rssi Zones Device

*This model accepts additional fields of type interface{}.*

## Structure

`StatsRssiZonesDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceId` | `*uuid.UUID` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_id": "0000188e-0000-0000-0000-000000000000",
  "rssi": 194,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

