
# Ap Search Wlan

*This model accepts additional fields of type interface{}.*

## Structure

`ApSearchWlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "00001fd4-0000-0000-0000-000000000000",
  "ssid": "ssid4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

