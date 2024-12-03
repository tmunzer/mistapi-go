
# Rssi Zone Device

*This model accepts additional fields of type interface{}.*

## Structure

`RssiZoneDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceId` | `uuid.UUID` | Required | - |
| `Rssi` | `int` | Required | RSSI threshold |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
  "rssi": 0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

