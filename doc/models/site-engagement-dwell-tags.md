
# Site Engagement Dwell Tags

add tags to visits within the duration (in seconds), available tags (passerby, bounce, engaged, stationed)

*This model accepts additional fields of type interface{}.*

## Structure

`SiteEngagementDwellTags`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `models.Optional[string]` | Optional | - |
| `Engaged` | `models.Optional[string]` | Optional | - |
| `Passerby` | `models.Optional[string]` | Optional | - |
| `Stationed` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "engaged": "300-14400",
  "stationed": "14400-43200",
  "bounce": "bounce2",
  "passerby": "passerby4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

