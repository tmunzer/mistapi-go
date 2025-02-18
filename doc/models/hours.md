
# Hours

Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun)

*This model accepts additional fields of type interface{}.*

## Structure

`Hours`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Fri` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Mon` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Sat` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Sun` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Thu` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Tue` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Wed` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "fri": "09:00-17:00",
  "mon": "09:00-17:00",
  "sat": "09:00-17:00",
  "sun": "09:00-17:00",
  "thu": "09:00-17:00",
  "tue": "09:00-17:00",
  "wed": "09:00-17:00",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

