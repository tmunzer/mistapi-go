
# Site Engagement Dwell Tag Names

name associated to each tag

*This model accepts additional fields of type interface{}.*

## Structure

`SiteEngagementDwellTagNames`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `*string` | Optional | **Default**: `"Visitor"` |
| `Engaged` | `*string` | Optional | **Default**: `"Associates"` |
| `Passerby` | `*string` | Optional | **Default**: `"Passerby"` |
| `Stationed` | `*string` | Optional | **Default**: `"Assets"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bounce": "Bounce",
  "engaged": "Engaged",
  "passerby": "Passer By",
  "stationed": "Stationed",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

