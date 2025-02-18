
# Response Stats Calls Summary

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseStatsCallsSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BadMinutesClient` | `*float64` | Optional | - |
| `BadMinutesSiteWan` | `*float64` | Optional | - |
| `BadMinutesWireless` | `*float64` | Optional | - |
| `NumAps` | `*int` | Optional | - |
| `NumUsers` | `*int` | Optional | - |
| `TotalMinutes` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bad_minutes_client": 526.0,
  "bad_minutes_site_wan": 3612.0,
  "bad_minutes_wireless": 1428.0,
  "num_aps": 1,
  "num_users": 3,
  "total_minutes": 5566,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

