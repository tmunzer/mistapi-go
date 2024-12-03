
# Pma Dashboard

*This model accepts additional fields of type interface{}.*

## Structure

`PmaDashboard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | description of the dashboard |
| `Label` | `*string` | Optional | group label name |
| `Name` | `*string` | Optional | name of the dashboard |
| `Url` | `*string` | Optional | url to access dashboard. Url will redirect the user to the dashboard |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "description": "Dashboard 1 description",
  "label": "Wireless",
  "name": "dashboard_1",
  "url": "https://api.mist.com/api/v1/forward/looker?jwt=...",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

