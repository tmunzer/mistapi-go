
# Service Packet

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePacket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceData` | `*string` | Optional | ata from service data |
| `ServiceUuid` | `*string` | Optional | UUID from service data |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "service_data": "service_data8",
  "service_uuid": "service_uuid4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

