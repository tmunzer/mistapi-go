
# Hours

hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).

**Note**: If the dow is not defined then it\u2019\ s treated as 00:00-23:59.

*This model accepts additional fields of type interface{}.*

## Structure

`Hours`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Fri` | `*string` | Optional | - |
| `Mon` | `*string` | Optional | - |
| `Sat` | `*string` | Optional | - |
| `Sun` | `*string` | Optional | - |
| `Thu` | `*string` | Optional | - |
| `Tue` | `*string` | Optional | - |
| `Wed` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "fri": "09:00-17:00",
  "mon": "09:00-17:00",
  "sat": "09:00-12:00",
  "sun": "09:00-12:00",
  "thu": "09:00-17:00",
  "tue": "09:00-17:00",
  "wed": "09:00-17:00",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

