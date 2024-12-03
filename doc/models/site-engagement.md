
# Site Engagement

**Note**: if hours does not exist, it’s treated as everyday of the week, 00:00-23:59. Currently we don’t allow multiple ranges for the same day

**Note**: default values for `dwell_tags`: passerby (1,300) bounce (301, 14400) engaged (14401, 28800) stationed (28801, 42000)

**Note**: default values for `dwell_tag_names`: passerby = “Passerby”, bounce = “Visitor”, engaged = “Associates”, stationed = “Assets”

*This model accepts additional fields of type interface{}.*

## Structure

`SiteEngagement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DwellTagNames` | [`*models.SiteEngagementDwellTagNames`](../../doc/models/site-engagement-dwell-tag-names.md) | Optional | - |
| `DwellTags` | [`*models.SiteEngagementDwellTags`](../../doc/models/site-engagement-dwell-tags.md) | Optional | add tags to visits within the duration (in seconds), available tags (passerby, bounce, engaged, stationed) |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).<br><br>**Note**: If the dow is not defined then it\u2019\ s treated as 00:00-23:59. |
| `MaxDwell` | `*int` | Optional | max time, default is 43200(12h), max is 68400 (18h)<br>**Default**: `43200`<br>**Constraints**: `>= 1`, `<= 68400` |
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

