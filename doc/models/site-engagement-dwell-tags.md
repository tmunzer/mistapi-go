
# Site Engagement Dwell Tags

add tags to visits within the duration (in seconds)

*This model accepts additional fields of type interface{}.*

## Structure

`SiteEngagementDwellTags`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `models.Optional[string]` | Optional | **Default**: `"301-14400"` |
| `Engaged` | `models.Optional[string]` | Optional | **Default**: `"14401-28800"` |
| `Passerby` | `models.Optional[string]` | Optional | **Default**: `"1-300"` |
| `Stationed` | `models.Optional[string]` | Optional | **Default**: `"28801-42000"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bounce": "301-14400",
  "engaged": "14401-28800",
  "passerby": "1-300",
  "stationed": "28801-42000",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

