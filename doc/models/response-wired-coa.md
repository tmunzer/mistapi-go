
# Response Wired Coa

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseWiredCoa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `Session` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_mac": "5c5b35000002",
  "port_id": "ge-0/0/0",
  "session": "0a2a11b8-4b30-40d8-a6d1-e91ea540d86f",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

