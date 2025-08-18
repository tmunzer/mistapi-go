
# Issued Client Certificate

*This model accepts additional fields of type interface{}.*

## Structure

`IssuedClientCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CommonName` | `*string` | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceId` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `SerialNumber` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "common_name": "john@corp.com",
  "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
  "serial_number": "13 00 13 03 23 EE D5 84 01",
  "created_time": 105.92,
  "modified_time": 229.04,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

