
# Alarms

*This model accepts additional fields of type interface{}.*

## Structure

`Alarms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmIds` | `[]uuid.UUID` | Required | - |
| `Note` | `*string` | Optional | Some text note describing the intent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alarm_ids": [
    "00001a02-0000-0000-0000-000000000000",
    "00001a03-0000-0000-0000-000000000000",
    "00001a04-0000-0000-0000-000000000000"
  ],
  "note": "note4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

