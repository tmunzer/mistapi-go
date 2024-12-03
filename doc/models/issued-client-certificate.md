
# Issued Client Certificate

*This model accepts additional fields of type interface{}.*

## Structure

`IssuedClientCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `DeviceId` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `SerialNumber` | `*string` | Optional | - |
| `SsoNameId` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
  "serial_number": "13 00 13 03 23 EE D5 84 01",
  "sso_name_id": "john@corp.com",
  "created_time": 105.92,
  "modified_time": 229.04,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

