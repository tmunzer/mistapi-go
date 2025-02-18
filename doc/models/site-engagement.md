
# Site Engagement

**Note**: if hours does not exist, it's treated as everyday of the week, 00:00-23:59. Currently we don't allow multiple ranges for the same day

*This model accepts additional fields of type interface{}.*

## Structure

`SiteEngagement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DwellTagNames` | [`*models.SiteEngagementDwellTagNames`](../../doc/models/site-engagement-dwell-tag-names.md) | Optional | Name associated to each tag |
| `DwellTags` | [`*models.SiteEngagementDwellTags`](../../doc/models/site-engagement-dwell-tags.md) | Optional | add tags to visits within the duration (in seconds) |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun) |
| `MaxDwell` | `*int` | Optional | Max time, default is 43200(12h), max is 68400 (18h)<br>**Default**: `43200`<br>**Constraints**: `>= 1`, `<= 68400` |
| `MinDwell` | `*int` | Optional | min time<br>**Constraints**: `>= 0` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "max_dwell": 43200,
  "dwell_tag_names": {
    "bounce": "bounce0",
    "engaged": "engaged2",
    "passerby": "passerby6",
    "stationed": "stationed4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "dwell_tags": {
    "bounce": "bounce0",
    "engaged": "engaged2",
    "passerby": "passerby6",
    "stationed": "stationed6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "hours": {
    "fri": "fri2",
    "mon": "mon8",
    "sat": "sat0",
    "sun": "sun6",
    "thu": "thu6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "min_dwell": 36,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

