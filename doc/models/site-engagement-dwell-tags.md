
# Site Engagement Dwell Tags

add tags to visits within the duration (in seconds), available tags (passerby, bounce, engaged, stationed)

## Structure

`SiteEngagementDwellTags`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `models.Optional[string]` | Optional | - |
| `Engaged` | `models.Optional[string]` | Optional | - |
| `Passerby` | `models.Optional[string]` | Optional | - |
| `Stationed` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "engaged": "300-14400",
  "stationed": "14400-43200",
  "bounce": "bounce2",
  "passerby": "passerby4"
}
```

